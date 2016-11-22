package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		templates := []string{
			filepath.Join("presentation_layer", "template", "menu.html"),
			filepath.Join("presentation_layer", "template", "footer.html"),
			filepath.Join("presentation_layer", "template", "roomScheduler.html"),
			filepath.Join("presentation_layer", "template", "app.html"),
			filepath.Join("presentation_layer", "template", "currentReservations.html"),
			filepath.Join("presentation_layer", "template", "createReservation.html"),
			filepath.Join("presentation_layer", "template", "modifyModal.html"),
			filepath.Join("presentation_layer", "template", "header.html"),
			filepath.Join("presentation_layer", "template", "waitList.html"),
		}
		t := (template.Must(template.ParseFiles(templates...)))
		t.ExecuteTemplate(rw, "app", nil)
	}
	if req.Method == "GET" {
		http.Redirect(rw, req, "/home", 303)
	}
}
