package model

import (
	"log"
	"os"
	"os/exec"
)

func GoFmt(f string) {
	_, err := exec.Command("go", "fmt", f).Output()
	if err != nil {
		log.Fatalf("go format err %s", err)
	}
}

func GoImport(f string) {
	result, err := exec.Command("goimports", f).Output()
	if err != nil {
		log.Fatalf("go import err %s", err)
	}
	file, err := os.Create(f)
	if err != nil {
		log.Fatalf("create file err %s", err)
	}
	_, _ = file.Write(result)
}
