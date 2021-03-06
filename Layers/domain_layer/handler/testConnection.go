package handler

import (
	"fmt"
	"net/http"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/data_source_layer/dB"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

func TestDb(rw http.ResponseWriter, req *http.Request) {
	db := dB.GetConnection()

	rows, _ := db.Query("SELECT * FROM userTable WHERE studentId=$1;", 1111111)
	var userId int
	var pass string
	for rows.Next() {
		rows.Scan(&userId, &pass)
	}
	defer dB.CloseConnection(db)
	fmt.Println(classes.User{userId, pass})
	rw.Write([]byte("Check terminal for result"))
}
