package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

func ReturnJson(rw http.ResponseWriter, req *http.Request) {
	user := classes.User{27192223, "NAME", "PASS"}

	rw.WriteHeader(http.StatusOK)

	err := json.NewEncoder(rw).Encode(user)
	if err != nil {
		fmt.Println("error message")
	}
}
