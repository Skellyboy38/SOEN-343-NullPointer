package mappers

import (
	"fmt"
	"time"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
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

	user, err := userMapper.GetById(userId)
	if err != nil {
		return err
	}
	newReservation := classes.Reservation{0, roomId, user, startTime, endTime}
	UOWSingleTon.RegisterNewReservation(newReservation)
	UOWSingleTon.Commit()
	return nil
}

func (reservationMapper *ReservationMapper) Update(reservationId, roomId, userId int, newStart, newEnd time.Time) error {
	reservations, err := reservationMapper.GetByRoomAndUserId(roomId, userId)
	if err != nil {
		return err
	}
	var reservation classes.Reservation
	for _, e := range reservations {
		if e.ReservationId == reservationId {
			reservation = e
		}
	}

	delete(reservationMapper.reservations, reservationId)

	reservation.StartTime = newStart
	reservation.EndTime = newEnd
	reservationMapper.reservations[reservationId] = reservation

	for i, e := range reservationMapper.reservationsByRoomId[roomId] {
		if e.ReservationId == reservationId {
			reservationMapper.reservationsByRoomId[roomId] = append(
				reservationMapper.reservationsByRoomId[roomId][:i],
				reservationMapper.reservationsByRoomId[roomId][i+1:]...)
			reservationMapper.reservationsByRoomId[roomId] = append(
				reservationMapper.reservationsByRoomId[roomId],
				reservation)
		}
	}

	for i, e := range reservationMapper.reservationsByUserId[userId] {
		if e.ReservationId == reservationId {
			reservationMapper.reservationsByUserId[userId] = append(
				reservationMapper.reservationsByUserId[userId][:i],
				reservationMapper.reservationsByUserId[userId][i+1:]...)
			reservationMapper.reservationsByUserId[userId] = append(
				reservationMapper.reservationsByUserId[userId],
				reservation)
		}
	}

	UOWSingleTon.RegisterDirtyReservations(reservation)
	if err := UOWSingleTon.Commit(); err != nil {
		return err
	}
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

func (reservationMapper *ReservationMapper) FilterOutUser(reservations []classes.Reservation, userId int) []classes.Reservation {
	otherReservations := []classes.Reservation{}

	for _, e := range reservations {
		if userId == e.User.StudentId {
			continue
		}
		otherReservations = append(otherReservations, e)
	}

	return otherReservations
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
	reservation := reservationMapper.reservations[id]
	delete(reservationMapper.reservationsByRoomId, reservation.Room)
	delete(reservationMapper.reservationsByUserId, reservation.User.StudentId)
	delete(reservationMapper.reservations, id)
	UOWSingleTon.RegisterDeleteReservation(id)
	err := UOWSingleTon.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (reservationMapper *ReservationMapper) SaveDeleted(reservationArray []int) error {
	if err := reservationMapper.reservationTDG.Delete(reservationArray); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (reservationMapper *ReservationMapper) SaveNew(reservationArray []classes.Reservation) error {
	for _, r := range reservationArray {
		reservationid, err := reservationMapper.reservationTDG.Create(r.Room, r.User.StudentId, r.StartTime, r.EndTime)
		if err != nil {
			fmt.Printf(" saveNew has a problem %v : \n", err)
			return err
			continue
		}
		r.ReservationId = reservationid
		reservationMapper.reservationsByRoomId.add(r.Room, []classes.Reservation{r})
		reservationMapper.reservationsByUserId.add(r.User.StudentId, []classes.Reservation{r})
	}
	reservationMapper.reservations.add(reservationArray)
	return nil
}

func (reservationMapper *ReservationMapper) SaveDirty(reservationArray []classes.Reservation) error {
	reservationIds := []int{}
	startTimes := []time.Time{}
	endTimes := []time.Time{}

	for _, e := range reservationArray {
		reservationIds = append(reservationIds, e.ReservationId)
		startTimes = append(startTimes, e.StartTime)
		endTimes = append(endTimes, e.EndTime)
	}
	if err := reservationMapper.reservationTDG.Update(reservationIds, startTimes, endTimes); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
