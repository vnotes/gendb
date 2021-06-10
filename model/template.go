package model

import (
	_ "embed"
	"html/template"
	"log"
	"os"
	"strings"
)

//go:embed db.tmpl
var b []byte

//go:embed db2.tmpl
var b2 []byte

func GetTmpl() []byte {
	if os.Getenv("type") == "2" {
		return b2
	}
	return b
}

func NewTemplateFile() *template.Template {
	tpl, err := template.New("db-template.tmpl").Funcs(
		template.FuncMap{
			"snakeCase2UpperCamelCase": snakeCase2UpperCamelCase,
			"snakeCase2LowerCamelCase": snakeCase2LowerCamelCase,
			"receiverName":             receiverName,
			"stringToLower":            stringToLower,
		},
	).Parse(string(GetTmpl()))
	if err != nil {
		log.Fatal("NewTemplateFile Error", err)
	}
	return tpl
}

func snakeCase2UpperCamelCase(s string) string {
	return strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, "_", " ")), " ", "")
}

func snakeCase2LowerCamelCase(s string) string {
	t := strings.ReplaceAll(strings.Title(strings.ReplaceAll(s, "_", " ")), " ", "")
	return strings.ToLower(s[0:1]) + t[1:]
}

func stringToLower(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", ""))
}

func receiverName(s string) string {
	return strings.ToLower(s[0:1])
}
