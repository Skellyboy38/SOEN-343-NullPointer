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

func (tdg UserTdg) Update(user classes.User) {
	DB.Exec("UPDATE userTable set studentId = '$1',  password = '$2' WHERE studentId = '$3';",
		user.StudentId, user.Password, user.StudentId)
}

func (tdg *UserTdg) GetByIdAndPass(id int, password string) (int, string, error) {
	studentIdString := strconv.Itoa(id)
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId='$1' and password='$2'",
		studentIdString, password)
	if err != nil {
		fmt.Println(err)
	}

	var studentId int

	if rows.Next() != true {
		return studentId, password, errors.New("No User Found")
	}

	err = rows.Scan(&studentId, &password)

	if err != nil {
		return studentId, password, nil
	} else {
		return studentId, password, err
	}
}

func (tdg UserTdg) Create(user classes.User) {
	fmt.Println(user)
	studentIdString := strconv.Itoa(user.StudentId)
	_, err := DB.Exec("INSERT INTO usertable (studentId, password) VALUES ('$1','$2');",
		studentIdString, user.Password)
	fmt.Println(err)
}
