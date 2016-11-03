package classes

import ()
import "time"

type Reservation struct {
	ReservationId int
	Room          Room
	User          User
	StartTime     time.Time
	EndTime       time.Time
}
