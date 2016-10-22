package dB

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func Init(driverName, dataSourceName string) {
	db, err := sql.Open(driverName, dataSourceName)
	Db = db
	if err != nil {
		fmt.Printf("***Connection to database failed at $1 \n", dataSourceName)
		log.Fatal("Could Not Connect To database")
	} else {
		fmt.Println("Connected to database")
	}
}
