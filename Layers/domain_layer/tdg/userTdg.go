package tdg

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

type UserTdg struct {
	AbstractTdg AbstractTDG
}

func (tdg UserTdg) Update(users []classes.User) error {

	for x, _ := range users {
		err := tdg.updateEach(users[x])
		if err != nil {
			return errors.New("Could not Update all")
		}
	}
	return nil
}

func (tdg UserTdg) updateEach(user classes.User) error {
	_, err := DB.Exec("UPDATE userTable set studentId = $1,  password = '$2' WHERE studentId = $3;",
		user.StudentId, user.Password, user.StudentId)
	return err
}

func (tdg *UserTdg) GetById(id int) (int, string, error) {
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId=$1;",
		id)

	if err != nil {
		fmt.Println(err)
	}

	var studentId int
	var foundPass string

	if rows.Next() == false {
		return studentId, foundPass, errors.New("User not found in Db")
	}

	err = rows.Scan(&studentId, &foundPass)
	return studentId, foundPass, err
}

func (tdg *UserTdg) GetByIdAndPass(id int, password string) (int, string, error) {
	rows, err := DB.Query("SELECT * FROM userTable WHERE studentId=$1;",
		id)

	if err != nil {
		fmt.Println(err)
	}

	var studentId int
	var foundPass string

	if rows.Next() == false {
		return studentId, password, errors.New("No User Found")
	}

	err = rows.Scan(&studentId, &foundPass)
	fmt.Println("Found studentId:" + strconv.Itoa(studentId))
	fmt.Println("Found password:" + foundPass)
	if err != nil {
		return studentId, password, err
	} else {
		if password == foundPass {
			fmt.Println("Found User in db")
			return studentId, foundPass, nil
		}
		fmt.Println("WrongPassword")
		return studentId, password, errors.New("WrongPassword")
	}
}

func (tdg UserTdg) Create(users []classes.User) error {
	for x, _ := range users {
		err := tdg.createEach(users[x])
		if err != nil {
			return errors.New("Could not Create all")
		} else {

		}
	}
	return nil
}

func (tdg UserTdg) createEach(user classes.User) error {
	_, err := DB.Exec("INSERT INTO usertable (studentId, password) VALUES ($1,'$2');",
		user.StudentId, user.Password)
	fmt.Println(err)
	return err
}

func (tdg UserTdg) deleteEach(userId int) error {
	_, err := DB.Exec("DELETE FROM userTable WHERE studentId = $1 ;", userId)
	return err
}

func (tdg UserTdg) Delete(userIds []int) error {
	for x, _ := range userIds {
		err := tdg.deleteEach(userIds[x])
		if err != nil {
			return errors.New("Could not delete all")
		}
	}
	return nil
}
