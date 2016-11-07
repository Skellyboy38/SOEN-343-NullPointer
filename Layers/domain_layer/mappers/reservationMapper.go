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
type ReservationMapper struct {
	reservations reservationIdentityMap
	reservationsByRoomId reservationByRoomIdBucketTable
	reservationTDG tdg.ReservationTDG
}

func InitReservationMapper() *ReservationMapper{
	return &ReservationMapper{make(map[int]classes.Reservation), map[int][]classes.Reservation{},tdg.ReservationTDG{}}
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
		roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByRoom(id)
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

			reservations = append(reservations,currentReservation)
		}
		reservationMapper.reservationsByRoomId.add(id,reservations)
		reservationMapper.reservations.add(reservations)

		return reservations, nil
	}
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
