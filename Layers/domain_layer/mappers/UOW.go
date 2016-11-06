package mappers

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
)

type userQueue []classes.User

var UOWSingleTon *UOW

type UOW struct {
	registeredNewUsers     userQueue
	registeredDirtyUsers   userQueue
	registeredDeletedUsers userQueue
	userMapper             *UserMapper
}

func InitUOW() {
	UOWSingleTon = &UOW{[]classes.User{}, []classes.User{}, []classes.User{},MapperBundle.UserMapper}
}

func (uow *UOW) RegisterNewUser(object classes.User) {
	uow.registeredNewUsers = append(uow.registeredNewUsers, object)
	fmt.Println(uow.registeredNewUsers)
}

func (uow *UOW) RegisterDirtyUser(object classes.User) {
	uow.registeredDirtyUsers = append(uow.registeredDirtyUsers, object)
}

func (uow *UOW) Commit() {
	fmt.Println("GOT TO COMMIT")
	fmt.Println(uow.registeredNewUsers)

	processedRegisteredNewUsers     := reverseUsers(reduceUserQueue(reverseUsers(uow.registeredNewUsers)))
	processedRegisteredDirtyUsers   := reverseUsers(reduceUserQueue(reverseUsers(uow.registeredDirtyUsers)))
	processedRegisteredDeletedUsers := convertToUserIdSlice(
		reverseUsers(
			reduceUserQueue(
				reverseUsers(
					uow.registeredDeletedUsers))))
	MapperBundle.UserMapper.SaveDeletedUsers(processedRegisteredDeletedUsers)
	MapperBundle.UserMapper.SaveDirtyUsers(processedRegisteredDirtyUsers)
	MapperBundle.UserMapper.SaveNewUsers(processedRegisteredNewUsers)

}

func reverseUsers(users []classes.User) []classes.User{
	reversedUsers := []classes.User{}
	for i := len(users)-1 ; i >=0; i--{
		reversedUsers = append(reversedUsers,users[i])
	}
	return reversedUsers
}

func reduceUserQueue(queue []classes.User) userQueue{
	reducedQueue := []classes.User{}
	exist := make(map[int]classes.User)
	for _,element := range queue{
		_, found := exist[element.StudentId]
		if found {
			continue
		}else{
			reducedQueue = append(reducedQueue,element)
		}
	}
	return reducedQueue
}

func convertToUserIdSlice(userSlice []classes.User) []int{
	intSlice := []int{}
	for _,x := range userSlice{
		intSlice =append(intSlice,x.StudentId)
	}
	return intSlice
}

