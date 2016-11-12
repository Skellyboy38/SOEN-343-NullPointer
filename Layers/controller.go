package main

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	fmt.Println("Application started")
	router := mux.NewRouter()
	err := startDb()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Db started")
	mappers.Init()
	mappers.InitUOW()
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./presentation_layer/js"))))
	router.HandleFunc("/login", handler.LoginGet).Methods("GET")
	router.HandleFunc("/login", handler.LoginForm).Methods("POST")
	router.HandleFunc("/createReservation", handler.CreateReservation).Methods("POST")
	router.HandleFunc("/home", handler.Home).Methods("GET")
	router.HandleFunc("/jsonexample", handler.ReturnJson).Methods("GET")
	router.HandleFunc("/reservations", handler.GetReservationsByRoomID).Methods("POST")
	router.HandleFunc("/testDbConnection", handler.TestDb).Methods("GET")
	router.HandleFunc("/testcookie", handler.TestCookie).Methods("GET")
	router.HandleFunc("/getcookie", handler.GetCookie).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func startDb() error {
	path, err := exec.LookPath("pg_ctl.exe")
	if err != nil {
		return err
	}
	cmd := exec.Command(path, "-D", "data_source_layer/setup/registry", "start")
	return cmd.Run()
}
