package main

import (
	"bytes"
	"log"
	"os"

	"github.com/vnotes/gendb/model"
	"github.com/vnotes/gendb/mysqldb"
)

func main() {
	var tName = mysqldb.Target
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
	var tGenName = tName + "_gen.go"
	file, err := os.Create(tGenName)
	if err != nil {
		log.Fatalf("create file err %s", err)
	}
	_, _ = file.WriteString(bf.String())
	_ = file.Close()
	model.GoFmt(tGenName)
	model.GoImport(tGenName)
}
