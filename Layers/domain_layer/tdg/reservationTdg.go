package tdg

import (
	"errors"
	"fmt"
	"time"
)

type ReservationTDG struct {
	AbstractTDG AbstractTDG
}

func (r *ReservationTDG) Create(roomId, userId []int, startTime, endTime []time.Time) []int {
	reservationIds := []int{}
	for i, _ := range roomId {
		res, err := DB.Exec("INSERT INTO reservation ( roomId, studentId, startTime, endTime)"+
			"VALUES ($1, '$2', $3, '$4');",
			roomId[i],
			userId[i],
			startTime[i],
			endTime[i])
		fmt.Println(err)
		id, err := res.LastInsertId()
		reservationIds = append(reservationIds, int(id))
	}
	return reservationIds
}

// stopped here because i need to have a get room and get user for when i
// do my select * from reservations
//func (r *ReservationTDG) ReadByRoom(id int) ([]int, []int, []time.Time, []time.Time, error) {
//	rows, err := DB.Query("SELECT * FROM reservation WHERE roomId=$1 ;",id)
//	if err != nil{
//		fmt.Println(err)
//	}
//	roomIds := []int{}
//	studentIds := []int{}
//	startTimes := []time.Time{}
//	endTimes   := []time.Time{}
//
//	var roomId int
//	var studentId int
//	var startTime  time.Time
//	var endTime time.Time
//
//
//	for rows.Next(){
//		err = rows.Scan(&id,&roomId,&studentId,&startTime,&endTime)
//		if err != nil{
//			return roomIds,studentIds,startTimes,endTimes,
//				errors.New("Could not Scan Reservation by RoomId")
//		}
//
//		roomIds    = append(roomIds,roomId)
//		studentIds = append(studentIds,studentId)
//		startTimes = append(startTimes,startTime)
//		endTimes   = append(endTimes,endTime)
//	}
//
//	return roomIds,studentIds,startTimes,endTimes, nil
//}

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

func (r *ReservationTDG) Delete(reservationId int) error {
	_, err := DB.Exec("DELETE * FROM reservation WHERE reservationId=$1;", reservationId)
	if err != nil {
		fmt.Println(err)
	}
	return err
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
