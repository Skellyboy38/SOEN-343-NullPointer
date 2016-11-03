package tdg

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"strconv"
	"fmt"
)

type RoomTDG struct {
	AbstractTDG AbstractTDG
}

func (r *RoomTDG) Create(room classes.Room) {
	_ , err :=	DB.Exec("INSERT INTO room (roomId, roomNumber)" + 
		"VALUES ('"+ strconv.Itoa(room.RoomId) + "," + room.RoomNumber + "');")
    fmt.Println(err)
}

func (r *RoomTDG) ReadAllRooms() {

}

func (r *RoomTDG) ReadAllAvailable() {

}

func (r *RoomTDG) Update() {

}

func (r *RoomTDG) Delete() {

}
