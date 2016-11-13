package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	//"time"
	//"os/user"	user := classes.User{studentId,password}
	"time"
)

func LoginGet(rw http.ResponseWriter, req *http.Request) {
	login := filepath.Join("presentation_layer", "template", "login.html")
	base := filepath.Join("presentation_layer", "template", "base.html")
	header := filepath.Join("presentation_layer", "template", "header.html")
	t := (template.Must(template.ParseFiles(base, login, header)))
	t.ExecuteTemplate(rw, "base", nil)
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
		rw.Write([]byte("Invalid id and password"))
		return
	}
	expire := time.Now().Add(time.Hour * 45)
	studentIdCookie := strconv.Itoa(verifiedUser.StudentId)
	studentIdAndName := "studentId=" + studentIdCookie
	cookie := http.Cookie{"studentId", studentIdCookie, "/", "localhost", expire, expire.Format(time.UnixDate), 86000, false, true, studentIdAndName, []string{studentIdAndName}}
	req.AddCookie(&cookie)
	http.SetCookie(rw, &cookie)
	http.Redirect(rw, req, "/home", 303)
	//user, err := mappers.MapperBundle.UserMapper.Get(studentId, password)
	//userMapper.Create(user.StudentId, password)
	//userMapper.Commit()
	fmt.Println(req.Form)
	fmt.Println(studentId)
	fmt.Println(password)
}
