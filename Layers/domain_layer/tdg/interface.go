package tdg

import (
	"database/sql"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
)

type AbstractTdg struct {
}

var DB *sql.DB

type TDG interface {
	Create()
	Read()
	Update()
	Delete()
}

func (tdg *AbstractTdg) GetConnection() {
	DB = dB.GetConnection()
}

func (tdg *AbstractTdg) CloseConnection() {
	dB.CloseConnection(DB)
}
