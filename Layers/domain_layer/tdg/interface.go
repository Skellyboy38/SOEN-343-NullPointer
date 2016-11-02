package tdg

import (
	"database/sql"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
)

type AbstractTDG struct {

}

type ReservationTDG struct {

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

func (r *ReservationTDG) Create() {

}

func (r *ReservationTDG) Read() {

}

func (r *ReservationTDG) Update() {

}

func (r *ReservationTDG) Delete() {

}