package dB

import (
	"database/sql"
	//"fmt"
	//"log"
	_ "github.com/lib/pq"
	//"fmt"
)

//var Db *sql.DB
var driverName = "postgres"
var dataSourceName = "user=soen343 sslmode=disable dbname=registry"
//
//func Init(driverName, dataSourceName string) {
//	db, err := sql.Open(driverName, dataSourceName)
//	Db = db
//	if err != nil {
//		fmt.Printf("***Connection to database failed at %s \n", dataSourceName)
//		log.Fatal("Could Not Connect To database")
//	} else {
//		fmt.Println("Connected to database")
//	}
//}

func GetConnection() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		//fmt.Printf("***Connection to database failed at %s \n", dataSourceName)
		//log.Fatal("Could Not Connect To database")
	} else {
		//fmt.Println("Connected to database")
		return db
	}
	return nil
}

func CloseConnection(conn *sql.DB) {
	conn.Close()
}