package classes

import (
// "fmt"
)

type Reservation struct {
	ReservationId int
	Room          Room
	User          User
	Time          TimeSlot
}
