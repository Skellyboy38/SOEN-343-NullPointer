package main

import (
	"fmt"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/classes"
	"github.com/gorilla/mux"
	//"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
	"log"
	"net/http"
	//_ "github.com/lib/pq"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
)

func main() {
	fmt.Println("Application started")
	 //dB.Init("postgres", "user=soen343 sslmode=disable dbname=registry")
	 //defer dB.Db.Close()
	router := mux.NewRouter()
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./presentation_layer/js"))))
	router.HandleFunc("/login", handler.LoginGet).Methods("GET")
	router.HandleFunc("/login", handler.LoginForm).Methods("POST")
	router.HandleFunc("/jsonexample", handler.ReturnJson).Methods("GET")
	router.HandleFunc("/testDbConnection", handler.TestDb).Methods("GET")
	router.HandleFunc("/testcookie", handler.TestCookie).Methods("GET")
	router.HandleFunc("/getcookie", handler.GetCookie).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
