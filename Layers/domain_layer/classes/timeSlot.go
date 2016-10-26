package classes

import (
	// "fmt"
	"time"
)

type TimeSlot struct {
	TimeSlotId int
	StartTime  time.Time
	Endtime    time.Time
}
