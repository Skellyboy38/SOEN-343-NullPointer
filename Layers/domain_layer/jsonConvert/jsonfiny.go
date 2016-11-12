package jsonConvert

import (
	"encoding/json"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"time"
)

type JsonReservation struct {
	ReservationID int       `json:"reservationID"`
	RoomNumber    int       `json:"roomNumber"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

func ReservationsJson(reservations []classes.Reservation) ([]byte, error) {
	formatedReservations := []JsonReservation{}
	for _, i := range reservations {
		fmt.Printf("id: %d  room: %d startTime %s endTime %s\n", i.ReservationId, i.Room, i.StartTime, i.EndTime)
		formatedReservation := JsonReservation{i.ReservationId,
			i.Room,
			i.StartTime,
			i.EndTime}
		formatedReservations = append(formatedReservations, formatedReservation)
	}
	return json.Marshal(formatedReservations)
}
