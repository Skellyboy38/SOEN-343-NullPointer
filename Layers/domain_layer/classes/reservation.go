package classes

import "time"

type Reservation struct {
	ReservationId int
	Room          int
	User          User
	StartTime     time.Time
	EndTime       time.Time
}
