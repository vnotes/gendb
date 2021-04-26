package model

import (
	"html/template"
	"log"
	"strings"
)

func NewTemplateFile() *template.Template {
	tpl, err := template.New("db.tmpl").Funcs(
		template.FuncMap{
			"snakeCase2UpperCamelCase": snakeCase2UpperCamelCase,
			"snakeCase2LowerCamelCase": snakeCase2LowerCamelCase,
			"receiverName":             receiverName,
		},
	).ParseFiles("db.tmpl")
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

func receiverName(s string) string {
	return strings.ToLower(s[0:1])
}
