package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
)

func TestLoginForm(t *testing.T) {

	err := startDb()
	if err != nil {
		log.Fatal(err)
	}

	mappers.Init()
	mappers.InitUOW()

	cases := []struct {
		id, password string
	}{
		{"1111111", "1111111"},
		{"2222222", "2222222"},
		{"3333333", "3333333"},
		{"4444444", "4444444"},
		{"5555555", "5555555"},
		{"6666666", "6666666"},
		{"7777777", "7777777"},
		{"kjdnfgoi", "rgjshg"},
		{"", ""},
		{"124512", "345643346"},
	}

	for i, c := range cases {
		t.Log("case:" + c.id)

		data := url.Values{}
		data.Set("id", c.id)
		data.Add("password", c.password)

		req, err := http.NewRequest(
			http.MethodPost,
			"http://localhost:9000/login",
			bytes.NewBufferString(data.Encode()),
		)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			t.Fatalf("could not create request %v", err)
		}

		rw := httptest.NewRecorder()

		handler.LoginForm(rw, req)

		if strings.Contains(rw.Body.String(), "Wrong Credentials") {
			if i != 7 && i != 8 && i != 9 {
				t.Errorf("wrong credentials expected correct")
				continue
			}
			if rw.Code != http.StatusOK {
				t.Errorf("expected status %s; got %d", http.StatusOK, rw.Code)
			}
			continue
		}

		if rw.Code != http.StatusPermanentRedirect {
			t.Errorf("expected status %s; got %d", http.StatusPermanentRedirect, rw.Code)
		}

		if i != 7 && i != 8 && i != 9 {
			resp := rw.Result()
			cookies := resp.Cookies()
			cookie := cookies[0]

			if cookie.Value != c.id {
				t.Errorf("studentId not in cookie")
			}
		}
	}

}
