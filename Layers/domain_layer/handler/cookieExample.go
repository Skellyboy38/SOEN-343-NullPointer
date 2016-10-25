package handler

import (
	"fmt"
	"net/http"
	"time"
)

func TestCookie(rw http.ResponseWriter, req *http.Request) {
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println(timeNow.Add(time.Second * 45))
	expire := time.Now().Add(time.Second * 45)
	cookie := http.Cookie{"test", "tcookie", "/", "localhost", expire, expire.Format(time.UnixDate), 86000, false, true, "test=tcookie", []string{"test=tcookie"}}
	http.SetCookie(rw, &cookie)
}
