package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl/part.gohtml", "tpl/parent.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "parent.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
