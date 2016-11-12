package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"net/http"
	"strconv"
)

func GetReservationsByRoomID(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()
	roomID, err := strconv.Atoi(req.FormValue("roomID"))
	fmt.Println(roomID)
	reservationsMapper := mappers.MapperBundle.ReservationMapper
	reservations, err := reservationsMapper.GetByRoomId(roomID)
	fmt.Println(reservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	jsonReservations, err := json.Marshal(reservations)
	rw.Write(jsonReservations)
}

func CreateReservation(rw http.ResponseWriter, req *http.Request) {
	//abstractTdg := mappers.MapperBundle.ReservationMapper.ReservationTdg.AbstractTdg

	//abstractTdg.GetConnection()
	//defer abstractTdg.CloseConnection()
	req.ParseForm()
	date := req.FormValue("date")
	room := req.FormValue("room")
	startTime := req.FormValue("start_time")
	endTime := req.FormValue("end_time")

	reservationMapper := mappers.MapperBundle.ReservationMapper

	reservationMapper.AddReservation(1111111, date, room, startTime, endTime)
	http.Redirect(rw, req, "/home", 303)
}
