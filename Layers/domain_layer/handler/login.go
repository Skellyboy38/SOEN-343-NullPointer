package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {
	if cookie, noCookie := req.Cookie("sessionId"); noCookie == nil {
		cookieIntVal, _ := strconv.Atoi(cookie.Value)
		if mappers.MapperBundle.SessionMapper.InMemoryBySessionId(cookieIntVal) {
			http.Redirect(rw, req, "/jsonexample", 200) // Change to main calendar page
		}
	}
	login := filepath.Join("presentation_layer", "template", "login.html")
	base := filepath.Join("presentation_layer", "template", "base.html")
	t := (template.Must(template.ParseFiles(base, login)))
	t.ExecuteTemplate(rw, "base", nil)
}

func LoginForm(rw http.ResponseWriter, req *http.Request) {
	mappers.MapperBundle.SessionMapper.SessionTdg.AbstractTDG.GetConnection()
	defer mappers.MapperBundle.SessionMapper.SessionTdg.AbstractTDG.CloseConnection()
	req.ParseForm()
	studentId, _ := strconv.Atoi(req.FormValue("id"))
	password := req.FormValue("password")
	user, err := mappers.MapperBundle.UserMapper.Get(studentId, password)
	if err == nil {
		currentSession, _ := mappers.MapperBundle.SessionMapper.Get(user)
		expire := time.Now().Add(time.Minute * 10)
		cookie := http.Cookie{"sessionId", strconv.Itoa(currentSession.SessionId), "/", "localhost", expire, expire.Format(time.UnixDate), 86000, false, true, "sessionId=" + strconv.Itoa(currentSession.SessionId), []string{"sessionId=" + strconv.Itoa(currentSession.SessionId)}}
		http.SetCookie(rw, &cookie)
	} else { // incorrect name and password

	}

	fmt.Println(req.Form)
	fmt.Println(studentId)
	fmt.Println(password)

}
