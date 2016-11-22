package jsonConvert

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

type JsonReservation struct {
	ReservationID int       `json:"reservationID"`
	StudentId     int       `json:"studentID"`
	RoomNumber    int       `json:"roomNumber"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

type JsonWaitingReservation struct {
	WaitlistID int       `json:"reservationID"`
	StudentId  int       `json:"studentID"`
	RoomNumber int       `json:"roomNumber"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

func ReservationsJson(reservations []classes.Reservation) ([]byte, error) {
	formatedReservations := []JsonReservation{}
	for _, i := range reservations {
		fmt.Printf("id: %d  room: %d startTime %s endTime %s\n", i.ReservationId, i.Room, i.StartTime, i.EndTime)
		formatedReservation := JsonReservation{i.ReservationId,
			i.User.StudentId,
			i.Room,
			i.StartTime,
			i.EndTime}
		formatedReservations = append(formatedReservations, formatedReservation)
	}
	return json.Marshal(formatedReservations)
}

func WaitListReservationsJson(waitList []classes.WaitlistReservation) ([]byte, error) {
	formatedWaitList := []JsonWaitingReservation{}
	for _, i := range waitList {
		fmt.Println("converting student id to json ",i.User.StudentId)
		formatedWaitingReservation := JsonWaitingReservation{i.WaitlistId,
			i.User.StudentId,
			i.Room,
			i.StartTime,
			i.EndTime}
		formatedWaitList = append(formatedWaitList, formatedWaitingReservation)
	}
	return json.Marshal(formatedWaitList)
}

func MessageJson(message string) ([]byte, error) {
	return json.Marshal(message)
}
