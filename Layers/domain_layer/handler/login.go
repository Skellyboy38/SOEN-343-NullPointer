package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {

	t := template.New("login") // can be anything
	t = template.Must(template.ParseFiles(filepath.Join("presentation_layer", "template", "login.html")))
	t.Execute(rw, "The sly red fox jumped over the fence") // the second inout can be any type of data, this is whet is passed to the html page or "template"
}

func LoginForm(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println(req.Form["username"])
	fmt.Println(req.Form["password"])
}
