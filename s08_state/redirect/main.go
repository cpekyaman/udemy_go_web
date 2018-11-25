package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barform", barform)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You requested foo with", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You called bar with", r.Method)
	// handle form post

	// redirect with explicit headers
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)

	// use http.Redirect
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func barform(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
