package mappers

import(

)
import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"errors"
	"time"
)

type reservationIdentityMap            map[int]classes.Reservation
type reservationByRoomIdBucketTable    map[int][]classes.Reservation
type reservationByStudentIdBucketTable    map[int][]classes.Reservation
type ReservationMapper struct {
	reservations reservationIdentityMap
	reservationsByRoomId reservationByRoomIdBucketTable
	reservationsByStudentId reservationByStudentIdBucketTable
	reservationTDG tdg.ReservationTDG
}

func InitReservationMapper() *ReservationMapper{
	return &ReservationMapper{make(map[int]classes.Reservation), map[int][]classes.Reservation{}, map[int][]classes.Reservation{},tdg.ReservationTDG{}}
}

func (reservationMapper *ReservationMapper) Create(roomId, userId []int, startTime, endTime []time.Time ) error{
	userMapper := MapperBundle.UserMapper
	user, err := userMapper.GetById(userId[0])
	if err != nil{
		return err
	}
	reservationIds := reservationMapper.reservationTDG.Create(roomId,userId,startTime,endTime)
	reservations := []classes.Reservation{}
	for i,_ := range roomId{
		reservations = append(reservations,classes.Reservation{reservationIds[i],
			roomId[i],
			user,
			startTime[i],
			endTime[i]})
	}
	reservationMapper.reservations.add(reservations)
	return nil
}


func (reservationMapper *ReservationMapper) GetByRoomId(id int) ([]classes.Reservation, error) {
	if reservationMapper.InMemoryByRoomId(id) {
		return reservationMapper.reservationsByRoomId[id], nil
	} else {
		_, roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByRoom(id)
		if err != nil {
			return []classes.Reservation{}, errors.New("No Reservations for that room doesnt exist")
		}
		reservations := []classes.Reservation{}

		for i, _ := range roomIds{
			student,err := MapperBundle.UserMapper.GetById(studentIds[i])
			if err != nil {
				return []classes.Reservation{}, errors.New("No Reservations for that room doesnt exist")
			}
			currentReservation :=classes.Reservation{id,roomIds[i],student,startTimes[i],endTimes[i]}

			reservations = append(reservations,currentReservation) //TODO why do we add here
		}
		reservationMapper.reservationsByRoomId.add(id,reservations)
		reservationMapper.reservations.add(reservations) //TODO and here

		return reservations, nil
	}
}

func (reservationMapper *ReservationMapper) GetByStudentId(id int) ([]classes.Reservation, error) {
	if reservationMapper.InMemoryStudentById(id) { //check if student's reservations are in memory
		return reservationMapper.reservationsByStudentId[id], nil
	} else {
		reservationIds, roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByUser(id)
		if err != nil {
			return []classes.Reservation{}, errors.New("This student does not have any reservations")
		}
		reservations := []classes.Reservation{}
		student,err := MapperBundle.UserMapper.GetById(studentIds[0])
		for i, _ := range reservationIds{
			if err != nil {
				return []classes.Reservation{}, errors.New("This student does not have any reservations")
			}
			currentReservation :=classes.Reservation{reservationIds[i],roomIds[i],student,startTimes[i],endTimes[i]}

			reservations = append(reservations,currentReservation) //mabye use set instead, possible repeats. 
		}
		reservationMapper.reservationsByStudentId.add(id,reservations)
		reservationMapper.reservations.add(reservations)

		return reservations, nil
	}
}

func (bucketTable reservationByStudentIdBucketTable) add(id int ,reservations []classes.Reservation){
	bucketTable[id] = append(bucketTable[id],reservations...)
}

func (bucketTable reservationByRoomIdBucketTable) add(id int ,reservations []classes.Reservation){
	bucketTable[id] = append(bucketTable[id],reservations...)
}

func (reservationMap reservationIdentityMap) add(reservations []classes.Reservation){ 
	for _ , e := range reservations{
		reservationMap[e.ReservationId] = e
	}
}

func (reservationMapper *ReservationMapper) InMemoryByRoomId(id int) bool {
	_, ok := reservationMapper.reservationsByRoomId[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (reservationMapper *ReservationMapper) InMemoryStudentById(id int) bool {
	_, ok := reservationMapper.reservationsByStudentId[id]
	if ok {
		return true
	} else {
		return false
	}
}