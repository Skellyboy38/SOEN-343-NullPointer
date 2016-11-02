package handler

import (
    "html/template"
    "net/http"
    "path/filepath"
	"fmt"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	app := filepath.Join("presentation_layer", "template", "app.html")
	t := (template.Must(template.ParseFiles(app)))
	err := t.ExecuteTemplate(rw, "app", nil)
	fmt.Print(err)
}