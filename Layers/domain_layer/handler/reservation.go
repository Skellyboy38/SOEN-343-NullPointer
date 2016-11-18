package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/jsonConvert"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"net/http"
	"strconv"
	"time"
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
	reservations, err := reservationsMapper.GetByRoomAndUserId(roomID, userID)

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

	roomIdint, _ := strconv.Atoi(roomId)
	userIDint, _ := strconv.Atoi(userId)
	startTimeformated, _ := time.Parse("yyyy-mm-dd hh:mm:ss", startTime)
	endTimeformated, _ := time.Parse("yyyy-mm-dd hh:mm:ss", endTime)
	reservationMapper := mappers.MapperBundle.ReservationMapper
	reservationMapper.Create(roomIdint, userIDint, startTimeformated, endTimeformated)
}

// func DeleteReservation(rw http.ResponseWriter, req *http.Request) {
// 	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
// 	abstractTdg.GetConnection()
// 	defer abstractTdg.CloseConnection()
// 	defer req.Body.Close()
// 	req.ParseForm()
// 	reservationID, err := strconv.Atoi(req.FormValue("reservationID"))

// 	reservationsMapper := mappers.MapperBundle.ReservationMapper

// 	if err := reservationsMapper.Delete(reservationID); err != nil {
// 		rw.WriteHeader(http.StatusExpectationFailed)
// 		bytes, _ := jsonConvert.MessageJson("Failure")
// 		rw.Write(bytes)
// 		return
// 	}

// 	rw.WriteHeader(http.StatusOK)
// 	bytes, _ := jsonConvert.MessageJson("Success")
// 	rw.Write(bytes)
// }
