package main

import (
	"os"
	"text/template"
)

func main() {
	tpl, _ := template.ParseFiles("tpl/slices.gohtml")

	turtles := []string{"leo", "mike", "don", "ralph"}
	tpl.Execute(os.Stdout, turtles)
}
