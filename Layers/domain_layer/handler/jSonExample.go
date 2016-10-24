package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/classes"
	"net/http"
)

func ReturnJson(rw http.ResponseWriter, req *http.Request) {
	user := classes.User{27192223}

	rw.WriteHeader(http.StatusOK)

	err := json.NewEncoder(rw).Encode(user)
	if err != nil {
		fmt.Println("error message")
	}
}
