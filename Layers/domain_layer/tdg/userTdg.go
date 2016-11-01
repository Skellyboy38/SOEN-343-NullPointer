package tdg

import (
	"database/sql"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

type UserTdg struct {
}

func (tdg *UserTdg) Update(dbConn *sql.DB, user classes.User) {
	dbConn.Exec("UPDATE userTable set")
}

func (tdg *UserTdg) GetByIdAndPass(dbConn *sql.DB, id int, password string) (int, string, error) {
	rows, err := dbConn.Query("SELECT * FROM userTable WHERE studentId='" + id + "' and password='" + password + "'")
	if err != nil {
		fmt.Println(err)
	}
	var studentId int

	for rows.Next() {
		err = rows.Scan(&studentId, &password)
	}
	if err != nil {
		return studentId, password, nil
	} else {
		return studentId, password, err
	}
}
