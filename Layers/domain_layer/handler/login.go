package handler

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {
	login := filepath.Join("presentation_layer", "template", "login.html")
	base := filepath.Join("presentation_layer", "template", "base.html")
	header := filepath.Join("presentation_layer", "template", "header.html")
	t := (template.Must(template.ParseFiles(base, login, header)))
	t.ExecuteTemplate(rw, "base", nil)
}

type ErrorMessage struct {
	Message template.HTML
}

func LoginForm(rw http.ResponseWriter, req *http.Request) {

	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg

	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()

	req.ParseForm()
	studentId, _ := strconv.Atoi(req.FormValue("id"))
	password := req.FormValue("password")

	userMapper := mappers.MapperBundle.UserMapper

	verifiedUser, err := userMapper.Get(studentId, password)

	if err != nil { // return error
		login := filepath.Join("presentation_layer", "template", "login.html")
		base := filepath.Join("presentation_layer", "template", "base.html")
		header := filepath.Join("presentation_layer", "template", "header.html")
		t := (template.Must(template.ParseFiles(base, login, header)))
		varmap := map[string]interface{}{
			"message": "Wrong Credentials",
		}
		t.ExecuteTemplate(rw, "base", varmap)

		return
	}
	expire := time.Now().Add(time.Hour * 45)
	studentIdCookie := strconv.Itoa(verifiedUser.StudentId)
	studentIdAndName := "studentId=" + studentIdCookie
	cookie := http.Cookie{"studentId", studentIdCookie, "/", "localhost", expire, expire.Format(time.UnixDate), 86000, false, false, studentIdAndName, []string{studentIdAndName}}
	req.AddCookie(&cookie)
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, req, "/home", 308)
}
