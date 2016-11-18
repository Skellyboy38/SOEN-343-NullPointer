package mappers

import ()
import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"time"
)

type reservationIdentityMap map[int]classes.Reservation
type reservationByRoomIdBucketTable map[int][]classes.Reservation
type reservationByUserIdBucketTable map[int][]classes.Reservation

type ReservationMapper struct {
	reservations         reservationIdentityMap
	reservationsByRoomId reservationByRoomIdBucketTable
	reservationsByUserId reservationByUserIdBucketTable
	reservationTDG       tdg.ReservationTDG
}

func InitReservationMapper() *ReservationMapper {
	return &ReservationMapper{
		make(map[int]classes.Reservation),
		map[int][]classes.Reservation{},
		map[int][]classes.Reservation{},
		tdg.ReservationTDG{}}
}

func (reservationMapper *ReservationMapper) Create(roomId, userId int, startTime, endTime time.Time) error {
	userMapper := MapperBundle.UserMapper
	reservationTDG := reservationMapper.reservationTDGe
	user, err := userMapper.GetById(userId[0])
	if err != nil {
		return err
	}
	newReservation := classes.Reservations{0, roomId, userId, startTime, endTime}
	UOWSingleTon.registeredNewReservations(newReservation)
	reservationIds := reservationMapper.reservationTDG.Create(roomId, userId, startTime, endTime)
	reservations := []classes.Reservation{}
	for i, _ := range roomId {
		reservations = append(reservations, classes.Reservation{reservationIds[i],
			roomId[i],
			user,
			startTime[i],
			endTime[i]})
	}
	reservationMapper.reservations.add(reservations)
	return nil
}

func (reservationMapper *ReservationMapper) GetByRoomAndUserId(roomId, userId int) ([]classes.Reservation, error) {
	if reservationMapper.InMemoryByUserId(userId) {
		return reservationMapper.reservationsByUserId[userId], nil
	} else {
		reservationIds, roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByUser(roomId, userId)
		if err != nil {
			return []classes.Reservation{}, err
		}
		reservations := []classes.Reservation{}

		for i, _ := range roomIds {
			student, err := MapperBundle.UserMapper.GetById(studentIds[i])
			if err != nil {
				return []classes.Reservation{}, err
			}
			currentReservation := classes.Reservation{reservationIds[i], roomIds[i], student, startTimes[i], endTimes[i]}

			reservations = append(reservations, currentReservation)
		}
		reservationMapper.reservationsByRoomId.add(roomId, reservations)
		reservationMapper.reservationsByUserId.add(userId, reservations)
		reservationMapper.reservations.add(reservations)

		return reservations, nil
	}
}

func (reservationMapper *ReservationMapper) GetByRoomId(roomId int) ([]classes.Reservation, error) {
	if reservationMapper.InMemoryByRoomId(roomId) {
		return reservationMapper.reservationsByRoomId[roomId], nil
	} else {
		reservationIds, roomIds, studentIds, startTimes, endTimes, err := reservationMapper.reservationTDG.ReadByRoom(roomId)
		if err != nil {
			return []classes.Reservation{}, err
		}
		reservations := []classes.Reservation{}

		for i, _ := range roomIds {
			student, err := MapperBundle.UserMapper.GetById(studentIds[i])
			if err != nil {
				return []classes.Reservation{}, err
			}
			currentReservation := classes.Reservation{reservationIds[i], roomIds[i], student, startTimes[i], endTimes[i]}

			reservations = append(reservations, currentReservation)
		}
		reservationMapper.reservationsByRoomId.add(roomId, reservations)
		reservationMapper.reservations.add(reservations)

		for _, e := range reservations {
			reservationMapper.reservationsByUserId[e.User.StudentId] = append(
				reservationMapper.reservationsByUserId[e.User.StudentId],
				e)
		}

		return reservations, nil
	}
}

func (bucketTable reservationByRoomIdBucketTable) add(id int, reservations []classes.Reservation) {
	bucketTable[id] = append(bucketTable[id], reservations...)
}

func (bucketTable reservationByUserIdBucketTable) add(id int, reservations []classes.Reservation) {
	bucketTable[id] = append(bucketTable[id], reservations...)
}

func (reservationMapper *ReservationMapper) AddReservation(id int, date string, room string, startTime string, endTime string) {
	//reservation := classes.Reservation{1, id, room, date, startTime, endTime}

	fmt.Println(id)
	fmt.Println(date)
	fmt.Println(room)
	fmt.Println(startTime)
	fmt.Println(endTime)
}

func (reservationMap reservationIdentityMap) add(reservations []classes.Reservation) {
	for _, e := range reservations {
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

func (reservationMapper *ReservationMapper) InMemoryByUserId(id int) bool {
	_, ok := reservationMapper.reservationsByUserId[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (reservationMapper *ReservationMapper) Delete(id int) error {
	delete(reservationMapper.reservations, id)
	err := reservationMapper.reservationTDG.Delete(id)
	return err
}
