package tdg

import (
	"database/sql"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
)

type AbstractTDG struct {
}

var DB *sql.DB

type TDG interface {
	Create()
	Read()
	Update()
	Delete()
}

func (tdg *AbstractTDG) GetConnection() {
	DB = dB.GetConnection()
}

func (tdg *AbstractTDG) CloseConnection() {
	dB.CloseConnection(DB)
}
