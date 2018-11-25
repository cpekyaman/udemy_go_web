package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("nf")
		if err != nil {
			handleUploadError(w, err)
			return
		}

		defer f.Close()

		ext := strings.Split(h.Filename, ".")[1]
		hash := sha1.New()
		io.Copy(hash, f)

		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		workDir, err := os.Getwd()
		if err != nil {
			handleUploadError(w, err)
			return
		}

		path := filepath.Join(workDir, "public", "images", fname)
		newFile, err := os.Create(path)
		if err != nil {
			handleUploadError(w, err)
			return
		}
		defer newFile.Close()

		f.Seek(0, 0)
		io.Copy(newFile, f)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func handleUploadError(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
