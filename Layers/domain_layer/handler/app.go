package handler

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func Home(rw http.ResponseWriter, req *http.Request) {
    templates := []string{
        filepath.Join("presentation_layer", "template", "menu.html"),
        filepath.Join("presentation_layer", "template", "footer.html"),
        filepath.Join("presentation_layer", "template", "app.html"),
    }
	t := (template.Must(template.ParseFiles(templates...)))
	t.ExecuteTemplate(rw, "app", nil)
}