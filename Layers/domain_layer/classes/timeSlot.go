package classes

import (
	// "fmt"
	"time"
)

type TimeSlot struct {
	TimeSlotId int
	start      time.Time
	end        time.Time
}
