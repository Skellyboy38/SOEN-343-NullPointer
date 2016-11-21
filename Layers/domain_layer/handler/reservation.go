package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/jsonConvert"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
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
	startTimeformated, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endTimeformated, _ := time.Parse("2006-01-02 15:04:05", endTime)
	reservationMapper := mappers.MapperBundle.ReservationMapper

	if err := reservationMapper.Create(roomIdint, userIDint, startTimeformated, endTimeformated); err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
	}

	rw.WriteHeader(http.StatusOK)
	bytes, _ := jsonConvert.MessageJson("Success")
	rw.Write(bytes)
}

func DeleteReservation(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()
	reservationID, _ := strconv.Atoi(req.FormValue("reservationID"))

	reservationsMapper := mappers.MapperBundle.ReservationMapper

	rw.Header().Set("Content-Type", "application/json")

	if err := reservationsMapper.Delete(reservationID); err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		bytes, _ := jsonConvert.MessageJson("Failure")
		rw.Write(bytes)
		return
	}

	rw.WriteHeader(http.StatusOK)
	bytes, _ := jsonConvert.MessageJson("Success")
	rw.Write(bytes)
}

func UpdateReservation(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()

	reservationID, _ := strconv.Atoi(req.FormValue("reservationID"))
	userID, _ := strconv.Atoi(req.FormValue("userID"))
	fmt.Printf("userID : %d", userID)
	roomID, _ := strconv.Atoi(req.FormValue("dataRoom"))
	newStart := req.FormValue("startTime")
	newEnd := req.FormValue("endTime")
	startTimeformated, _ := time.Parse("2006-01-02 15:04:05", newStart)
	endTimeformated, _ := time.Parse("2006-01-02 15:04:05", newEnd)
	fmt.Printf("startTimeFormated %v \n", startTimeformated)
	fmt.Printf("endTimeFormated %v \n", endTimeformated)
	rw.Header().Set("Content-Type", "application/json")
	reservationMapper := mappers.MapperBundle.ReservationMapper

	if err := reservationMapper.Update(reservationID, roomID, userID, startTimeformated, endTimeformated); err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		bytes, _ := jsonConvert.MessageJson("Failure")
		rw.Write(bytes)
		return
	}

	rw.WriteHeader(http.StatusOK)
	bytes, _ := jsonConvert.MessageJson("Success")
	rw.Write(bytes)

}

func GetReservationsOthers(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()

	roomID, err := strconv.Atoi(req.FormValue("dataRoom"))
	userID, err := strconv.Atoi(req.FormValue("UserID"))
	reservationsMapper := mappers.MapperBundle.ReservationMapper
	reservations, err := reservationsMapper.GetByRoomId(roomID)

	otherReservations := reservationsMapper.FilterOutUser(reservations, userID)

	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}

	jsonReservations, err := jsonConvert.ReservationsJson(otherReservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonReservations)

}
