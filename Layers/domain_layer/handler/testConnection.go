package handler

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/layers/data_source_layer/dB"
	"github.com/Skellyboy38/SOEN-343-NullPointer/layers/domain_layer/classes"
	"net/http"
)

func TestDb(rw http.ResponseWriter, req *http.Request) {
	db := dB.Db
	_, err := db.Exec("INSERT INTO user (studentId) VALUES ('1234') ;")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT EXISTS (SELECT 1 FROM accounts WHERE username=$1 LIMIT 1);", 1234)
	var userId int
	var name string
	var pass string
	for rows.Next() {
		err := rows.Scan(&userId, &name, &pass)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(classes.User{userId, name, pass})
}
