package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func MyHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "main.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("main.gohtml"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(MyHandler))
}
