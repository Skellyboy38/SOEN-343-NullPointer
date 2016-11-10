package handler

import (
	"net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"fmt"
	"strconv"
	"encoding/json"
)

func GetReservationsByRoomID(rw http.ResponseWriter, req *http.Request) {
	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg
	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
	defer req.Body.Close();
	req.ParseForm()
	roomID, err := strconv.Atoi(req.FormValue("roomID"))
	fmt.Println(roomID)
	reservationsMapper := mappers.MapperBundle.ReservationMapper
	reservations, err := reservationsMapper.GetByRoomId(roomID)
	fmt.Println(reservations)
	if err != nil{
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	jsonReservations , err := json.Marshal(reservations)
	rw.Write(jsonReservations)
}