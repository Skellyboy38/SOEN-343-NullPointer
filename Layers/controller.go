package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
)

func main() {
	fmt.Println("Application started")
	router := mux.NewRouter()
	mappers.Init()
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./presentation_layer/js"))))
	router.HandleFunc("/login", handler.LoginGet).Methods("GET")
	router.HandleFunc("/login", handler.LoginForm).Methods("POST")
	router.HandleFunc("/home", handler.Home).Methods("GET")
	router.HandleFunc("/jsonexample", handler.ReturnJson).Methods("GET")
	router.HandleFunc("/testDbConnection", handler.TestDb).Methods("GET")
	router.HandleFunc("/testcookie", handler.TestCookie).Methods("GET")
	router.HandleFunc("/getcookie", handler.GetCookie).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
