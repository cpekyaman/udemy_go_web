package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const form string = `
<form method="POST" action="/upload" enctype="multipart/form-data">
<input type="file" name="q" />
<input type="submit" />
</form>`

func main() {
	http.HandleFunc("/", renderForm)
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, form)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		fmt.Println("\nFile: ", f, "\nHeader: ", h, "\nErr: ", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(filepath.Join("./storage/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		io.WriteString(w, form+`<p>`+string(bs)+`</p>`)
	}
}
