package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Alert(rw http.ResponseWriter, req *http.Request) {
	t := template.New("alert")
	t = template.Must(template.ParseFiles(filepath.Join("js", "alert.js")))
	t.Execute(rw, nil)
}
