package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/img/copy", copyImage)
	http.HandleFunc("/img/serve", serveImage)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="archangel.jpg"></img>`)
}

func copyImage(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/archangel.jpg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/archangel.jpg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer f.Close()

	fs, err := f.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fs.ModTime(), f)
}
