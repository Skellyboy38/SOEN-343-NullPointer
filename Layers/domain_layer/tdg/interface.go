package tdg

import (
	"database/sql"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
)

type abstractTdg struct {
}
type TDG interface {
	Create()
	Remove()
	Update()
	Delete()
}

func (tdg *TDG) GetConnection() *sql.DB {
	dbConn := dB.GetConnection()
	return dbConn
}

func (tdg *TDG) CloseConnection(conn *sql.DB) {
	dB.CloseConnection(conn)
}
