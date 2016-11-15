package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/jsonConvert"
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
	roomID, err := strconv.Atoi(req.FormValue("dataRoom"))
	reservationsMapper := mappers.MapperBundle.ReservationMapper
	reservations, err := reservationsMapper.GetByRoomId(roomID)

	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}

	jsonReservations, err := jsonConvert.ReservationsJson(reservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonReservations)
}

func GetReservationsByUserID(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()
	roomID, err := strconv.Atoi(req.FormValue("dataRoom"))
	userID, err := strconv.Atoi(req.FormValue("userID"))

	reservationsMapper := mappers.MapperBundle.ReservationMapper
	reservations, err := reservationsMapper.GetByRoomAndUserId(roomID,userID)

	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}

	jsonReservations, err := jsonConvert.ReservationsJson(reservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonReservations)
}

func CreateReservation(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	req.ParseForm()
	roomId := req.FormValue("dataRoom")
	userId := req.FormValue("userID")
	startTime := req.FormValue("startTime")
	endTime := req.FormValue("endTime")

	reservationMapper := mappers.MapperBundle.ReservationMapper

	reservationMapper.AddReservation(1111111, date, room, startTime, endTime)
	http.Redirect(rw, req, "/home", 303)
}

func DeleteReservation(rw http.ResponseWriter, req *http.Request){
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()
	reservationID, err := strconv.Atoi(req.FormValue("reservationID"))

	reservationsMapper := mappers.MapperBundle.ReservationMapper

	if err := reservationsMapper.Delete(reservationID); err != nil{
		rw.WriteHeader(http.StatusExpectationFailed)
		bytes, _ := jsonConvert.MessageJson("Failure")
		rw.Write(bytes)
		return
	}

	rw.WriteHeader(http.StatusOK)
	bytes, _ := jsonConvert.MessageJson("Success")
	rw.Write(bytes)
}
