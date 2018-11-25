package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "cenk",
	})
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Your Cookie: ", c)
}
