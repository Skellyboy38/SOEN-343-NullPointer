package tdg

import (
	"errors"
	"fmt"
	"time"
)

type ReservationTDG struct {
	AbstractTDG AbstractTDG
}

func (r *ReservationTDG) ReadByRoom(roomId int) ([]int, []int, []int, []time.Time, []time.Time, error) {
	rows, err := DB.Query("SELECT * FROM reservation WHERE roomId=$1 ;", roomId)
	if err != nil {
		fmt.Println(err)
	}
	reservationIds := []int{}
	roomIds := []int{}
	studentIds := []int{}
	startTimes := []time.Time{}
	endTimes := []time.Time{}

	var reservationId int
	var studentId int
	var startTime time.Time
	var endTime time.Time

	for rows.Next() {
		err = rows.Scan(&reservationId, &roomId, &studentId, &startTime, &endTime)
		if err != nil {
			return reservationIds, roomIds, studentIds, startTimes, endTimes,
				errors.New("Could not Scan Reservation by RoomId")
		}
		reservationIds = append(reservationIds, reservationId)
		roomIds = append(roomIds, roomId)
		studentIds = append(studentIds, studentId)
		startTimes = append(startTimes, startTime)
		endTimes = append(endTimes, endTime)
	}

	return reservationIds, roomIds, studentIds, startTimes, endTimes, nil
}

func (r *ReservationTDG) Update() {

}

func (r *ReservationTDG) Delete(reservationIds []int) error {

	for _, i := range reservationIds {
		_, err := DB.Exec("DELETE * FROM reservation WHERE reservationId=$1;", i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (r *ReservationTDG) ReadByUser(roomId, userId int) ([]int, []int, []int, []time.Time, []time.Time, error) {
	rows, err := DB.Query("SELECT * FROM reservation WHERE roomId=$1 and studentId=$2 ;", roomId, userId)
	if err != nil {
		fmt.Println(err)
	}
	reservationIds := []int{}
	roomIds := []int{}
	studentIds := []int{}
	startTimes := []time.Time{}
	endTimes := []time.Time{}

	var reservationId int
	var studentId int
	var startTime time.Time
	var endTime time.Time

	for rows.Next() {
		err = rows.Scan(&reservationId, &roomId, &studentId, &startTime, &endTime)
		if err != nil {
			return reservationIds, roomIds, studentIds, startTimes, endTimes,
				errors.New("Could not Scan Reservation by RoomId")
		}
		reservationIds = append(reservationIds, reservationId)
		roomIds = append(roomIds, roomId)
		studentIds = append(studentIds, studentId)
		startTimes = append(startTimes, startTime)
		endTimes = append(endTimes, endTime)
	}

	return reservationIds, roomIds, studentIds, startTimes, endTimes, nil
}

func (r *ReservationTDG) Create(roomId, studentId int, startTime, endTime time.Time) (int, error) {
	reservationId, err := DB.Exec("INSERT into reservations (roomId, studentId, startTime, endTime) VALUES ($1,$2,$3,$4) ",
		roomId,
		studentId,
		startTime,
		endTime)

	if err != nil {
		return -1, errors.New("Could not create reservation")
	}
	resId, err := reservationId.LastInsertId()

	if err != nil {
		fmt.Printf("Cannot get last inserted id : %v", err)
		return -1, err

	}

	return int(resId), nil
}
