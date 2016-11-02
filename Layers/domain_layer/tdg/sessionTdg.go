package tdg

import (
	"errors"
	"strconv"
)

type SessionTdg struct {
	AbstractTdg AbstractTdg
}

func (tdg SessionTdg) Read(studentId int) (int, int, error) {
	dbConn := DB
	rows, _ := dbConn.Query("SELECT * FROM session WHERE studentId='" + strconv.Itoa(studentId) + "'")
	if rows.Next() == false {
		return 0, 0, errors.New("User has no session in the db")
	}
	var sessionId int
	rows.Scan(&sessionId, &studentId)
	return sessionId, studentId, nil
}

func (tdg SessionTdg) Create(studentId int) (int, error){
	dbConn := DB
	result,err := dbConn.Exec("INSERT INTO session VALUES ('"+strconv.Itoa(studentId)+"';")
	if err != nil{
		id , _ := result.LastInsertId()
		return int(id) , nil
	}else{
		return 0 , errors.New("Could not create a new session")
	}
}
