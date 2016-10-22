package main

import (
	"fmt"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/Modules/classes"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/Modules/dB"
	"github.com/gorilla/mux"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/Modules/mappers"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Application started")
	// dB.Init("postgres", "user=soen343 password=soen343 sslmode=disable dbname=registry")
	// defer dB.Db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/login", handler.LoginGet)
	router.HandleFunc("/js/alert.js", handler.Alert) /* seems like we need a route to the js files too
	which sucks. im sure there is a different way to do this and i found one but its crap */
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
