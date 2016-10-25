package main

import (
	"fmt"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/layers/data_source_layer/dB"
	"github.com/gorilla/mux"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/mappers"
	"github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Application started")
	dB.Init("postgres", "user=soen343 password=soen343 sslmode=disable dbname=registry")
	defer dB.Db.Close()
	router := mux.NewRouter()
	router.PathPrefix("/prpesentation_layer/js/").Handler(http.StripPrefix("/presentation_layer/js/", http.FileServer(http.Dir("./presentation_layer/js"))))
	router.HandleFunc("/login", handler.LoginGet).Methods("GET")
	router.HandleFunc("/login", handler.LoginForm).Methods("POST")
	router.HandleFunc("/jsonexample", handler.ReturnJson).Methods("GET")
	router.HandleFunc("/testDbConnection", handler.TestDb).Methods("GET")
	router.HandleFunc("/testcookie", handler.TestCookie).Methods("GET")
	router.HandleFunc("/getcookie", handler.GetCookie).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
