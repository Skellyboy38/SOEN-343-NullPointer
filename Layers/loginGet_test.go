package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
)

func TestLoginGet(t *testing.T) {

	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:9000/login",
		nil,
	)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rw := httptest.NewRecorder()
	handler.LoginGet(rw, req)

	if rw.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", rw.Code)
	}
}
