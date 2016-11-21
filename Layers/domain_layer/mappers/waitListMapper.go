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
	fmt.Println("in the mapper")
	if err != nil {
		return err
	}
	newWaitReservation := classes.WaitlistReservation{0, roomId, user, startTime, endTime}
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

func (waitListMapper *WaitListMapper) GetByRoomId(roomId int) ([]classes.WaitlistReservation, error){
	waitReservationIds, roomIds, studentIds, startTimes, endTimes, err := waitListMapper.waitListTDG.ReadByRoom(roomId)
		if err != nil {
			return []classes.WaitlistReservation{}, err
		}
		waitingReservations := []classes.WaitlistReservation{}

		for i, _ := range roomIds {
			student, err := MapperBundle.UserMapper.GetById(studentIds[i])
			if err != nil {
				return []classes.WaitlistReservation{}, err
			}
			currentwaitingReservation := classes.WaitlistReservation{waitReservationIds[i], roomIds[i], student, startTimes[i], endTimes[i]}
			waitingReservations = append(waitingReservations, currentwaitingReservation)
		}
		//reservationMapper.reservationsByRoomId.add(roomId, reservations)
		waitListMapper.waitList.add(waitingReservations)
		/*for _, e := range reservations {
			reservationMapper.reservationsByUserId[e.User.StudentId] = append(
				reservationMapper.reservationsByUserId[e.User.StudentId],
				e)
		}*/
		return waitingReservations, nil
}

func (waitListMapper *WaitListMapper) Delete(waitingReservationId int) error{
	//waitingreservation := waitListMapper.waitList[waitingReservationId]
	//delete(waitListMapper.reservationsByRoomId, reservation.Room)
	//delete(waitListMapper.reservationsByUserId, reservation.User.StudentId)
	delete(waitListMapper.waitList,waitingReservationId )
	//UOWSingleTon.RegisterDeleteWaitingReservation(id)
	err := UOWSingleTon.Commit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
	
}

func (waitListMapper *WaitListMapper) InMemoryByWaitingReservationId(id int) bool {
	_, ok := waitListMapper.waitList[id]
	if ok {
		return true
	} else {
		return false
	}
} 

//functions for the wait list identity map
func (waitListMap waitListIdentityMap) add(waitingList []classes.WaitlistReservation) {
	for _, e := range waitingList {
		waitListMap[e.WaitlistId] = e
	}
}

//end of functions for the wait list identity map


//GetByRoomId return all the waiting reservations with the room id.