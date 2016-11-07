package tdg

import (
	"fmt"
	"time"
	"errors"
)

type ReservationTDG struct {
	AbstractTDG AbstractTDG
}

func (r *ReservationTDG) Create(roomId, userId []int, startTime, endTime []time.Time) []int{
	reservationIds := []int{}
	for i,_ := range roomId{
		res, err := DB.Exec("INSERT INTO reservation ( roomId, studentId, startTime, endTime)"+
			"VALUES ($1, '$2', $3, '$4');",
			roomId[i],
			userId[i],
			startTime[i],
			endTime[i])
		fmt.Println(err)
		id , err := res.LastInsertId()
		reservationIds = append(reservationIds,int(id))
	}
	return reservationIds
}

// stopped here because i need to have a get room and get user for when i
// do my select * from reservations
func (r *ReservationTDG) ReadByRoom(id int) ([]int, []int, []time.Time, []time.Time, error) {
	rows, err := DB.Query("SELECT * FROM reservation WHERE roomId=$1 ;",id)
	if err != nil{
		fmt.Println(err)
	}
	roomIds := []int{}
	studentIds := []int{}
	startTimes := []time.Time{}
	endTimes   := []time.Time{}

	var roomId int
	var studentId int
	var startTime  time.Time
	var endTime time.Time

	if rows.Next() == false{
		return roomIds,studentIds,startTimes,endTimes,
			errors.New("Could not get Reservations by RoomId")
	}

	err = rows.Scan(&id,&roomId,&studentId,&startTime,&endTime)

	if err != nil{
		return roomIds,studentIds,startTimes,endTimes,
			errors.New("Could not Scan Reservation by RoomId")
	}
	roomIds    = append(roomIds,roomId)
	studentIds = append(studentIds,studentId)
	startTimes = append(startTimes,startTime)
	endTimes   = append(endTimes,endTime)

	for rows.Next(){
		if err != nil{
			return roomIds,studentIds,startTimes,endTimes,
				errors.New("Could not Scan Reservation by RoomId")
		}

		roomIds    = append(roomIds,roomId)
		studentIds = append(studentIds,studentId)
		startTimes = append(startTimes,startTime)
		endTimes   = append(endTimes,endTime)
	}

	return roomIds,studentIds,startTimes,endTimes, nil
}

func (r *ReservationTDG) ReadByUser() {

}

func (r *ReservationTDG) Update() {

}

func (r *ReservationTDG) Delete() {

}
