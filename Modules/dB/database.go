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

func Insert(query string) error {
	_, err := Db.Exec(query)
	return err
}

func GetAll(query string) ([][]string, error) {
	vals := [][]string{}
	gotten := []string{}
	rows, err := Db.Query(query)
	for rows.Next() {
		rows.Scan(&gotten)
		vals = append(vals, gotten)
	}

	// var found string
	// for rows.Next() {
	// 	rows.Scan(&found)
	// 	vals = append(vals, found)
	// }
	return vals, err
}
