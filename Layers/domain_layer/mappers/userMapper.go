package mappers

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"errors"
)

type UserMapper struct {
	users map[string]classes.User
	userTdg tdg.UserTdg
}

func InitUserMapper () *UserMapper{
	return &UserMapper{make(map[int]classes.User),tdg.UserTdg{}}
}

func (userMap *UserMapper)InMemory(username string) bool{
	_, ok:= userMap.users[username]
	if ok {
		return true
	}else{
		return false
	}
}

func (userMap *UserMapper) Get(username, password string) (classes.User, error){
	if userMap.InMemory(username){
		return userMap.users[username], nil
	}else{
		studentId, _, _, err:=userMap.userTdg.GetByUsernameAndPass(username, password)
		if err != nil {
			return classes.User{}, errors.New("User not in Memory")
		}
		return classes.User{studentId,username,password}, nil
	}
}
