package main

import (
	"os"
	"text/template"
)

func main() {
	tpl, _ := template.ParseFiles("tpl/wisdom.gohtml")
	tpl.Execute(os.Stdout, 42)
}
