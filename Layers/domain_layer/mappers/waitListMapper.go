package mappers

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"time"
	"fmt"

)

type waitListIdentityMap map[int]classes.WaitlistReservation

type WaitListMapper struct {
	waitList waitListIdentityMap
	waitListTDG tdg.WaitlistReservationTDG
}

func InitWaitListMapper() *WaitListMapper {
	return &WaitListMapper{make(map[int]classes.WaitlistReservation), tdg.WaitlistReservationTDG{}}
}

//create a reservation on the wait list
func (waitListMapper *WaitListMapper) Create(roomId, userId int, startTime, endTime time.Time) error {
	userMapper := MapperBundle.UserMapper
	user, err := userMapper.GetById(userId)
	if err != nil {
		return err
	}
	newWaitReservation := classes.WaitlistReservation{0, roomId, user, startTime, endTime}
	//waitListMapper.waitList = append(waitListMapper.waitList, newWaitReservation)
	UOWSingleTon.RegisterNewWaitingReservation(newWaitReservation)
	UOWSingleTon.Commit()
	return nil
}

func (waitListMapper *WaitListMapper) SaveNew(waitingReservationArray []classes.WaitlistReservation) error {
	for _, w := range waitingReservationArray {
		waitingReservationId, err := waitListMapper.waitListTDG.Create(w.Room, w.User.StudentId, w.StartTime, w.EndTime)
		if err != nil {
			fmt.Printf(" saveNew has a problem %v : \n", err)
			return err
			continue
		}
		w.WaitlistId = waitingReservationId
		//reservationMapper.reservationsByRoomId.add(r.Room, []classes.Reservation{r})
		//reservationMapper.reservationsByUserId.add(r.User.StudentId, []classes.Reservation{r})
	}
	waitListMapper.waitList.add(waitingReservationArray)
	return nil
}

func (waitListMap waitListIdentityMap) add(waitList []classes.WaitlistReservation) {
	for _, e := range waitList {
		waitListMap[e.WaitlistId] = e
	}
}

//GetByRoomId return all the waiting reservations with the room id.