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
	rows, err := DB.Query("SELECT * FROM waitlistMaster WHERE roomId=$1 ;", roomId)
	if err != nil {
		fmt.Println(err)
	}
	waitListreservationIds := []int{}
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
			return waitListreservationIds, roomIds, studentIds, startTimes, endTimes,
				errors.New("Could not Scan WaitListReservation by RoomId")
		}
		waitListreservationIds = append(waitListreservationIds, waitListreservationId)
		roomIds = append(roomIds, roomId)
		studentIds = append(studentIds, studentId)
		startTimes = append(startTimes, startTime)
		endTimes = append(endTimes, endTime)
	}

	return waitListreservationIds, roomIds, studentIds, startTimes, endTimes, nil
}

func (r *WaitlistReservationTDG) Delete(waitListreservationIds []int) error {

	for _, i := range waitListreservationIds {
		_, err := DB.Exec("DELETE FROM waitlistMaster WHERE waitlistID=$1;", i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (r *WaitlistReservationTDG) ReadByUser(roomId, userId int) ([]int, []int, []int, []time.Time, []time.Time, error) {
	rows, err := DB.Query("SELECT * FROM waitlistMaster WHERE roomId=$1 and studentId=$2 ;", roomId, userId)
	if err != nil {
		fmt.Println(err)
	}
	waitListreservationIds := []int{}
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
			return waitListreservationIds, roomIds, studentIds, startTimes, endTimes,
				errors.New("Could not Scan Master Wait list by RoomId")
		}
		waitListreservationIds = append(waitListreservationIds, waitListreservationId)
		roomIds = append(roomIds, roomId)
		studentIds = append(studentIds, studentId)
		startTimes = append(startTimes, startTime)
		endTimes = append(endTimes, endTime)
	}

	return waitListreservationIds, roomIds, studentIds, startTimes, endTimes, nil
}

func (r *WaitlistReservationTDG) Create(roomId, studentId int, startTime, endTime time.Time) (int, error) {
	waitListreservationId := 0
	res, err := DB.Query("INSERT INTO waitlistMaster (roomId, studentId, startTime, endTime) VALUES ($1,$2,$3,$4) RETURNING waitlistID;",
		roomId,
		studentId,
		startTime.Format("2006-01-02 15:04:05"),
		endTime.Format("2006-01-02 15:04:05"))

	res.Scan(&waitListreservationId)
	if err != nil {
		fmt.Printf(startTime.Format("2006-01-02 15:04:05"))
		fmt.Printf(endTime.Format("2006-01-02 15:04:05"))
		fmt.Println(err)
		return -1, errors.New("Could not create Wait list reservation")
	}

	if err != nil {
		fmt.Printf("Cannot get last inserted id : %v", err)
		return -1, err

	}

	return int(waitListreservationId), nil
}
