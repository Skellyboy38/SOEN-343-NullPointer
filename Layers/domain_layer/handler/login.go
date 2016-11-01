package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {
	if cookie, noCookie := req.Cookie("sessionId"); noCookie == nil {
		cookieIntVal, _ := strconv.Atoi(cookie.Value)
		if mappers.MapperBundle.SessionMapper.InMemory(cookieIntVal) {
			http.Redirect(rw, req, "/jsonexample", 200) // Change to main calendar page
		}
	}
	login := filepath.Join("presentation_layer", "template", "login.html")
	base := filepath.Join("presentation_layer", "template", "base.html")
	t := (template.Must(template.ParseFiles(base, login)))
	t.ExecuteTemplate(rw, "base", nil)
}

func LoginForm(rw http.ResponseWriter, req *http.Request) {
	db := dB.GetConnection()
	req.ParseForm()
	username := req.Form["username"]
	password := req.Form["password"]
	user, err := mappers.MapperBundle.UserMapper.Get(username, password)
	if err != nil {
		sessionId, studentId, err := tdg.SessionTdg{}.FindByStudentId(db, user.StudentId)
		if err != nil { // no session found, create one

		} else {
			newSession := classes.Session{sessionId, studentId}
			mappers.MapperBundle.SessionMapper.AddToMap(newSession)
		}
	} else { // incorrect name and password

	}

	fmt.Println(req.Form)
	fmt.Println(username)
	fmt.Println(password)

}
