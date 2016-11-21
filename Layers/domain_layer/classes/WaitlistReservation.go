package classes

import "time"

type WaitlistReservation struct {
	WaitlistId int
	Room       int
	User       User
	StartTime  time.Time
	EndTime    time.Time
}
