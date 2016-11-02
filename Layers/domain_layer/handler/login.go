package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
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
	mappers.MapperBundle.SessionMapper.SessionTdg.AbstractTdg.GetConnection()
	defer mappers.MapperBundle.SessionMapper.SessionTdg.AbstractTdg.CloseConnection()
	req.ParseForm()
	id , _:= strconv.Atoi(req.FormValue("id"))
	password := req.FormValue("password")
	user, err := mappers.MapperBundle.UserMapper.Get(id, password)
	if err != nil {
		sessionId, studentId, err := mappers.MapperBundle.SessionMapper.SessionTdg.Read(user.StudentId)
		if err != nil { // no session found, create one

		} else {
			student ,_ := mappers.MapperBundle.UserMapper.Get(studentId,password)
			newSession := classes.Session{sessionId, student}
			mappers.MapperBundle.SessionMapper.AddToMap(newSession)
		}
	} else { // incorrect name and password

	}

	fmt.Println(req.Form)
	fmt.Println(id)
	fmt.Println(password)

}
