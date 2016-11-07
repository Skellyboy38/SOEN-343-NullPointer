package handler

import (
    "net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
	"strconv"
	"encoding/json"
	"log"
	"io"
	"fmt"
)

type Message struct {
	RoomID string `json:"roomID"`
}

func ReservationsByRoom(rw http.ResponseWriter, req *http.Request){

	abstractTdg := mappers.MapperBundle.UserMapper.UserTdg.AbstractTdg

	abstractTdg.GetConnection()
	defer abstractTdg.CloseConnection()
    decoder := json.NewDecoder(req.Body)
    for {
        var m Message
		err := decoder.Decode(&m)
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
		roomID, err := strconv.Atoi(m.RoomID)
		fmt.Printf("%v", roomID)
		// reservationMapper := mappers.MapperBundle.ReservationMapper
		// reservations, err := reservationMapper.GetReservationsByRoomId(roomID)
		
		// if err != nil {
		// rw.WriteHeader(http.StatusExpectationFailed)
		// 	return
		// }
		// jsonReservations , err := json.Marshal(reservations)
		// rw.Header().Set("Content-Type", "application/json")
		// rw.Write(jsonReservations);
    }
}