package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/jsonConvert"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"net/http"
	"strconv"
	"time"
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
	roomIdint, _ := strconv.Atoi(roomId)
	userIDint, _ := strconv.Atoi(userId)

	startTimeformated, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endTimeformated, _ := time.Parse("2006-01-02 15:04:05", endTime)

	waitListMapper := mappers.MapperBundle.WaitListMapper
	waitListMapper.Create(roomIdint, userIDint, startTimeformated, endTimeformated)
}
