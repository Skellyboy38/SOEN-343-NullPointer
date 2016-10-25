package handler

import (
	"fmt"
	"net/http"
)

func GetCookie(rw http.ResponseWriter, req *http.Request) {
	cookie, _ := req.Cookie("test")
	fmt.Println(cookie.Value)
}
