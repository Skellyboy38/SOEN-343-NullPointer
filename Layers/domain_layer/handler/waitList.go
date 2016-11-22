package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/jsonConvert"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
)

func AddToWaitList(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()

	req.ParseForm()
	roomId := req.FormValue("dataRoom")
	userId := req.FormValue("userID")
	startTime := req.FormValue("startTime")
	endTime := req.FormValue("endTime")
	roomIdInt, _ := strconv.Atoi(roomId)
	userIDInt, _ := strconv.Atoi(userId)

	startTimeformated, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endTimeformated, _ := time.Parse("2006-01-02 15:04:05", endTime)

	waitListMapper := mappers.MapperBundle.WaitListMapper

	if err := waitListMapper.Create(roomIdInt, userIDInt, startTimeformated, endTimeformated); err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
	}

	rw.WriteHeader(http.StatusOK)
	bytes, _ := jsonConvert.MessageJson("Success")
	rw.Write(bytes)
}

func GetWaitListEntriesByRoom(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()

	req.ParseForm()
	roomId := req.FormValue("dataRoom")
	waitListMapper := mappers.MapperBundle.WaitListMapper
	roomIdInt, _ := strconv.Atoi(roomId)
	reservations, err := waitListMapper.GetByRoomId(roomIdInt)
	fmt.Println(reservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}

	jsonReservations, err := jsonConvert.WaitListReservationsJson(reservations)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonReservations)
}

func GetWaitListReservationsByUserID(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close()
	req.ParseForm()
	userID, err := strconv.Atoi(req.FormValue("userID"))
	waitListMapper := mappers.MapperBundle.WaitListMapper
	waitList, err := waitListMapper.GetByUserId(userID)

	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}

	jsonWaitList, err := jsonConvert.WaitListReservationsJson(waitList)
	if err != nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonWaitList)
}

func RemoveWaitListEntriesById(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()

	req.ParseForm()
	waitListId := req.FormValue("waitListId")

	waitListMapper := mappers.MapperBundle.WaitListMapper
	waitListIdInt, _ := strconv.Atoi(waitListId)
	waitListMapper.Delete(waitListIdInt)
}
