package dB

import (
	"database/sql"
	_"github.com/lib/pq"
)

var driverName = "postgres"
var dataSourceName = "user=soen343 sslmode=disable dbname=registry"

func GetConnection() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err == nil {
		return db
	}
	return nil
}

func CloseConnection(conn *sql.DB) {
	conn.Close()
}