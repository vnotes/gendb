package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vnotes/gendb/model"
	"github.com/vnotes/gendb/mysqldb"
)

func main() {
	target := os.Getenv("target")
	if target == "" {
		log.Fatalf("needed env table target. target=xxxoo go run main.go")
	}
	var tName = target
	table, err := mysqldb.NewDatabaseTableInfo(tName)
	if err != nil {
		log.Fatalf("get table info error %s", err)
	}
	tpl := model.NewTemplateFile()
	bf := &bytes.Buffer{}
	err = tpl.Execute(bf, table)
	if err != nil {
		log.Fatalf("exe tpl err %s", err)
		return
	}
	var targetDir = "model"
	if !isDirExist(targetDir) {
		if _, err = os.Create("model"); err != nil {
			log.Fatalf("create file err %s", err)
		}
	}
	if os.Getenv("type") == "2" {
		targetDir = targetDir + "/" + strings.ToLower(strings.ReplaceAll(target, "_", "")) + "api"
		if !isDirExist(targetDir) {
			if err = os.Mkdir(targetDir, os.ModePerm); err != nil {
				log.Fatalf("create file err %s", err)
			}
		}
	}
	var tGenName = tName + "_gen.go"
	file, err := os.Create(tGenName)
	if err != nil {
		log.Fatalf("create file err %s", err)
	}
	_, _ = file.WriteString(bf.String())
	_ = file.Close()
	model.GoFmt(tGenName)
	model.GoImport(tGenName)
	_ = os.Rename(tGenName, fmt.Sprintf("%s/%s", targetDir, tGenName))
}

func isDirExist(d string) bool {
	_, err := os.Stat(d)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
