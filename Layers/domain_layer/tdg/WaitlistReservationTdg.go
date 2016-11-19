package tdg

import (
	"errors"
	"fmt"
	"time"
)

type WaitlistReservationTDG struct {
	AbstractTDG AbstractTDG
}

func (r *WaitlistReservationTDG) ReadByRoom(roomId int) ([]int, []int, []int, []time.Time, []time.Time, error) {
	rows, err := DB.Query("SELECT * FROM waitlistReservation WHERE roomId=$1 ;", roomId)
	if err != nil {
		fmt.Println(err)
	}
	waitListwaitListreservationIds := []int{}
	roomIds := []int{}
	studentIds := []int{}
	startTimes := []time.Time{}
	endTimes := []time.Time{}

	var waitListreservationId int
	var studentId int
	var startTime time.Time
	var endTime time.Time

	for rows.Next() {
		err = rows.Scan(&waitListreservationId, &roomId, &studentId, &startTime, &endTime)
		if err != nil {
			return waitListwaitListreservationIds, roomIds, studentIds, startTimes, endTimes,
				errors.New("Could not Scan Reservation by RoomId")
		}
		waitListwaitListreservationIds = append(waitListwaitListreservationIds, waitListreservationId)
		roomIds = append(roomIds, roomId)
		studentIds = append(studentIds, studentId)
		startTimes = append(startTimes, startTime)
		endTimes = append(endTimes, endTime)
	}

	return waitListwaitListreservationIds, roomIds, studentIds, startTimes, endTimes, nil
}