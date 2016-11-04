package tdg

import (
	"errors"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

type UserTdg struct {
	AbstractTdg AbstractTDG
}

func (tdg UserTdg) Update(user classes.User) {
	DB.Exec("UPDATE userTable set studentId = $1,  password = '$2' WHERE studentId = $3;",
		user.StudentId, user.Password, user.StudentId)
}

func (tdg *UserTdg) GetByIdAndPass(id int, password string) (int, string, error) {
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId=$1 ;",
	id)
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

func (tdg UserTdg) Create(user classes.User) error {
	_, err := DB.Exec("INSERT INTO usertable (studentId, password) VALUES ($1,'$2');",
		user.StudentId, user.Password)
	return err
}

func (tdg UserTdg) DeleteEach(userId int) error {
	_, err := DB.Exec("DELETE FROM userTable WHERE studentId = $1 ;", userId)
	return err
}

func (tdg UserTdg) Delete(userIds []int) error {
	for x, _ := range userIds {
		err := tdg.DeleteEach(userIds[x])
		if err != nil {
			return errors.New("Could not delete all")
		}
	}
	return nil
}
