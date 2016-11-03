package classes

import ()
import "time"

type Reservation struct {
	ReservationId int
	Room          Room
	User          User
	StartTime     time.Duration
	EndTime       time.Duration
}
