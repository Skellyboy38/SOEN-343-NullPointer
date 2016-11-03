package tdg

import (
	"errors"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"strconv"
)

type UserTdg struct {
	AbstractTdg AbstractTDG
}

func (tdg *UserTdg) Update(user classes.User) {
	DB.Exec("UPDATE userTable set")
}

func (tdg *UserTdg) GetByIdAndPass(id int, password string) (int, string, error) {
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId='" + strconv.Itoa(id) + "' and password='" + password + "'")
	if err != nil {
		fmt.Println(err)
	}
	var studentId int

	if rows.Next() != true {
		return studentId, password, errors.New("No User Found")
	}
	for rows.Next() {
		err = rows.Scan(&studentId, &password)
	}
	if err != nil {
		return studentId, password, nil
	} else {
		return studentId, password, err
	}
}

func (tdg UserTdg) Create(user classes.User) {
	fmt.Println(user)
	_, err := DB.Exec("INSERT INTO usertable (studentId, password) VALUES ('" + strconv.Itoa(user.StudentId) + "','" + user.Password + "');")
	fmt.Println(err)
}
