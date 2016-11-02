package tdg

import (
	"errors"
	"strconv"
)

type SessionTdg struct {
	AbstractTdg AbstractTdg
}

func (tdg *SessionTdg) Read(studentId int) (int, int, error) {
	dbConn := DB
	rows, _ := dbConn.Query("SELECT * FROM session WHERE studentId='" + strconv.Itoa(studentId) + "'")
	if rows.Next() == false {
		return 0, 0, errors.New("User has no session in the db")
	}
	var sessionId int
	rows.Scan(&sessionId, &studentId)
	return sessionId, studentId, nil
}
