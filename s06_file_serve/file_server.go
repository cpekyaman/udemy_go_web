package main

import (
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
