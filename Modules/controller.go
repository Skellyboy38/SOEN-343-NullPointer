package main

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/dB"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/mapper"
	//	 	"github.com/Skellyboy38/SOEN-343-NullPointer/Modules/handler"
	"log"
	"net/http"
)

type Ball struct {
	name string
	size int
}

func main() {
	fmt.Println("Application started")
	ball := Ball{"puma", 12}
	mapper.Insert(&ball)
	dB.Init("postgres", "user=soen343 password=soen343 sslmode=disable dbname=registry")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
