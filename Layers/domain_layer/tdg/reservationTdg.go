package tdg

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"strconv"
	"fmt"
)

type ReservationTDG struct {
	AbstractTDG AbstractTDG
}

func (r *ReservationTDG) Create(reservation classes.Reservation) {
	_ , err :=	DB.Exec("INSERT INTO reservation (reservationId, roomId, studentId, startTime, endTime)" + 
		"VALUES ('" + strconv.Itoa(reservation.ReservationId) + "," + strconv.Itoa(reservation.Room.RoomId) + "," + strconv.Itoa(reservation.User.StudentId) + "," + reservation.StartTime.String() + "," + reservation.EndTime.String() + "');")
	fmt.Println(err)
}

func (r *ReservationTDG) ReadByRoom() {

}

func (r *ReservationTDG) ReadByUser() {

}

func (r *ReservationTDG) Update() {

}

func (r *ReservationTDG) Delete() {

}
