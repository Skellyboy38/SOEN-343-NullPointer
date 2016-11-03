package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {
	login := filepath.Join("presentation_layer", "template", "login.html")
	base := filepath.Join("presentation_layer", "template", "base.html")
	t := (template.Must(template.ParseFiles(base, login)))
	t.ExecuteTemplate(rw, "base", nil)
}

func LoginForm(rw http.ResponseWriter, req *http.Request) {
	mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg.GetConnection()
	defer mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg.CloseConnection()
	req.ParseForm()
	studentId, _ := strconv.Atoi(req.FormValue("id"))
	password := req.FormValue("password")
	//user, err := mappers.MapperBundle.UserMapper.Get(studentId, password)


	fmt.Println(req.Form)
	fmt.Println(studentId)
	fmt.Println(password)

}
