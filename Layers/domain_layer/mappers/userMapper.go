package mappers

import (
	"errors"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
)

type userIdentityMap map[int]classes.User

type UserMapper struct {
	users   userIdentityMap
	UserTdg tdg.UserTdg
}

func InitUserMapper() *UserMapper {
	return &UserMapper{make(map[int]classes.User), tdg.UserTdg{}}
}

func (userMap *UserMapper) InMemory(id int) bool {
	_, ok := userMap.users[id]
	if ok {
		return true
	} else {
		return false
	}
}

func (userMap *UserMapper) Get(id int, password string) (classes.User, error) {
	if userMap.InMemory(id) {
		return userMap.users[id], nil
	} else {
		studentId, _, err := userMap.UserTdg.GetByIdAndPass(id, password)
		if err != nil {
			return classes.User{}, errors.New("User not in Memory")
		}
		return classes.User{studentId, password}, nil
	}
}

func (userMap userIdentityMap) add(user classes.User) {
	userMap[user.StudentId] = user
}

func (userMapper *UserMapper) Create(studentId int,password string) (classes.User, error){
	if userMapper.InMemory(studentId){
		return classes.User{}, errors.New("already exists")
	}
	user := classes.User{studentId,password}
	userMapper.users.add(user)
	tdg.UOWSingleTon.RegisterNew(user)
	return user, nil
}
