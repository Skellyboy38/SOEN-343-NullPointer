package tdg

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"strconv"
)

type UserTdg struct {
	AbstractTdg AbstractTdg
}

func (tdg *UserTdg) Update( user classes.User) {
	DB.Exec("UPDATE userTable set")
}

func (tdg *UserTdg) GetByIdAndPass( id int, password string) (int, string, error) {
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId='" + strconv.Itoa(id) + "' and password='" + password + "'")
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
