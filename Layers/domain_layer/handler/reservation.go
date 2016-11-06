package handler

import (
    "net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"strconv"
	"encoding/json"
	"fmt"
)

func GetAllRoomReservations(rw http.ResponseWriter, req *http.Request) {
    // 
}

func GetUserRoomReservations(rw http.ResponseWriter, req *http.Request) {
    //
}

func ReservationByRoom(rw http.ResponseWriter, req *http.Request){

	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg

	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()

	roomId,err := strconv.Atoi(req.URL.Query().Get("roomId"))
	reservationMapper := mappers.MapperBundle.ReservationMapper

	reservations, err := reservationMapper.GetByRoomId(roomId)

	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
	jsonReservations , err := json.Marshal(reservations)

	rw.Write(jsonReservations)

}