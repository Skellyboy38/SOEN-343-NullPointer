package tdg

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"database/sql"
)

type UserTdg struct {

}

func (tdg *UserTdg) Update(dbConn *sql.DB, user classes.User){
	dbConn.Exec("UPDATE userTable set")
}

func (tdg *UserTdg) GetByUsernameAndPass(dbConn *sql.DB, username, password string) (int , string ,string ,error){
	rows , err := dbConn.Query("SELECT * FROM userTable WHERE username='"+username+"' and password='"+password+"'")
	if err != nil{
		fmt.Println(err)
	}
	var studentId int

	for rows.Next(){
		err = rows.Scan(&studentId,&username,&password)
	}
	if err != nil{
		return studentId, username, password, nil
	}else{
		return studentId, username, password, err
	}
}