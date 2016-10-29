package handler

import (
	"net/http"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

func TestDb(rw http.ResponseWriter, req *http.Request) {
	db := dB.GetConnection()

	rows, _ := db.Query("SELECT * FROM userTable WHERE studentId=$1;", 1111111)
	var userId int
	var name string
	var pass string
	for rows.Next() {
		rows.Scan(&userId, &name, &pass)
	}
	dB.CloseConnection(db)
	fmt.Println(classes.User{userId, name, pass})
	rw.Write([]byte("Check terminal for result"))
}
