package mappers

import(

)
import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"errors"
	"fmt"
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


func (reservationMapper *ReservationMapper) GetReservationsByRoomId(id int) ([]classes.Reservation, error) {
	if reservationMapper.InMemoryByRoomId(id) {
		return reservationMapper.reservationsByRoomId[id], nil
	} else {
		roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByRoom(id)
		if err != nil {
			return []classes.Reservation{}, errors.New(fmt.Sprintf("Reservations for room %v that room exist", id))
		}
		reservations := []classes.Reservation{}

		for i, _ := range roomIds{
			student, err := MapperBundle.UserMapper.GetUserById(studentIds[i])

			if err != nil {
				return []classes.Reservation{}, errors.New("No reservations for user.")
			}
			currentReservation := classes.Reservation{id,roomIds[i],student,startTimes[i],endTimes[i]}

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
