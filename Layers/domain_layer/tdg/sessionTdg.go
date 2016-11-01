package tdg

import (
	"database/sql"
	"errors"
	"strconv"
)

type SessionTdg struct {
}

func (tdg *SessionTdg) FindByStudentId(dbConn *sql.DB, studentId int) (int, int, error) {
	rows, _ := dbConn.Query("SELECT * FROM session WHERE studentId='" + strconv.Itoa(studentId) + "'")
	if rows.Next() == false {
		return 0, 0, errors.New("User has no session in the db")
	}
	var sessionId int
	rows.Scan(&sessionId, &studentId)
	return sessionId, studentId, nil
}
