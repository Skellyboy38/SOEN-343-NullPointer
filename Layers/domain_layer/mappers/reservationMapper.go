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
	// reservationTDG := reservationMapper.reservationTDG
	user, err := userMapper.GetById(userId)
	if err != nil {
		return err
	}
	newReservation := classes.Reservation{0, roomId, user, startTime, endTime}
	UOWSingleTon.RegisterNewReservation(newReservation)
	// reservations := []classes.Reservation{}
	// for i, _ := range roomId {
	// 	reservations = append(reservations, classes.Reservation{reservationIds[i],
	// 		roomId[i],
	// 		user,
	// 		startTime[i],
	// 		endTime[i]})
	// }
	// reservationMapper.reservations.add(reservations)
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

func (reservationMapper *ReservationMapper) InMemoryByReservationId(id int) bool {
	_, ok := reservationMapper.reservations[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (reservationMapper *ReservationMapper) Delete(id int) error {
	delete(reservationMapper.reservations, id)
	return nil
}

func (reservationMapper *ReservationMapper) SaveDeleted(reservationArray []int) {
	reservationMapper.reservationTDG.Delete(reservationArray)
}

func (reservationMapper *ReservationMapper) SaveNew(reservationArray []classes.Reservation) {
	for _, r := range reservationArray {
		reservationid, err := reservationMapper.reservationTDG.Create(r.Room, r.User.StudentId, r.StartTime, r.EndTime)
		if err != nil {
			fmt.Printf(" saveNew has a problem %v : \n", err)
			continue
		}
		r.ReservationId = reservationid
		reservationMapper.reservationsByRoomId.add(r.Room, []classes.Reservation{r})
		reservationMapper.reservationsByRoomId.add(r.User.StudentId, []classes.Reservation{r})
	}
	reservationMapper.reservations.add(reservationArray)
}

func (reservationMapper *ReservationMapper) SaveDirty(reservationArray []classes.Reservation) {
	// reservationMapper.reservationTDG.Update(reservationArray)
}
