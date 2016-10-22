package main

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/dB"
	// "github.com/Skellyboy38/SOEN-343-NullPointer/Modules/mappers"
	//	 	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Application started")

	dB.Init("postgres", "user=soen343 password=soen343 sslmode=disable dbname=registry")
	defer dB.Db.Close()

	log.Fatal(http.ListenAndServe(":9000", nil))

}
