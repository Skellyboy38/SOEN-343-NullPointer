package mappers

import (
	"errors"
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/tdg"
	"sync"
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
		student := userMap.users[id]
		if student.Password == password {
			return student, nil
		}
		return student, errors.New("WrongPassword")
	} else {
		_, _, err := userMap.UserTdg.GetByIdAndPass(id, password)
		if err != nil {
			return classes.User{}, errors.New("User doesnt exist")
		}
		foundUserInDb := classes.User{id, password}
		userMap.users.add(foundUserInDb)
		return foundUserInDb, nil
	}
}

func (userMap *UserMapper) GetById(id int) (classes.User, error) { // add tdg that searches by
	if userMap.InMemory(id) { // only id and finish with check to db
		return userMap.users[id], nil
	} else {
		studentId, password, err := userMap.UserTdg.GetById(id)
		if err != nil {
			return classes.User{}, err
		}
		student := classes.User{studentId, password}
		userMap.users.add(student)
		return student, nil
	}
}

func (userMap userIdentityMap) add(user classes.User) {
	userMap[user.StudentId] = user
}

func (userMapper *UserMapper) SaveNew(userArray []classes.User) error {
	return userMapper.UserTdg.Create(userArray)
}

func (userMapper *UserMapper) SaveDeleted(userArray []int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, e := range userArray {
			delete(userMapper.users, e)
		}
	}()

	go func() {
		defer wg.Done()
		tdg.UserTdg{}.Delete(userArray)
	}()

	wg.Wait()
}

func (userMapper *UserMapper) SaveDirty(users []classes.User) error {
	return tdg.UserTdg{}.Update(users)
}

func (userMapper *UserMapper) Create(studentId int, password string) (classes.User, error) {
	if userMapper.InMemory(studentId) {
		return classes.User{}, errors.New("already exists")
	}
	user := classes.User{studentId, password}
	fmt.Println(user)
	userMapper.users.add(user)
	UOWSingleTon.RegisterNewUser(user)
	return user, nil
}

func (userMapper *UserMapper) Commit() {
	UOWSingleTon.Commit()
}
