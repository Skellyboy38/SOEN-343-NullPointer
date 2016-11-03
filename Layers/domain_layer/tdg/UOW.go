package tdg

import (
	"fmt"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"reflect"
	"strings"
)

type user classes.User

type objectQueue []interface{}

var UOWSingleTon UOW

type UOW struct {
	registeredNew     objectQueue
	registeredDirty   objectQueue
	registeredDeleted objectQueue
}

func InitUOW() {
	UOWSingleTon = UOW{objectQueue{}, objectQueue{}, objectQueue{}}
}

func (uow *UOW) RegisterNew(object interface{}) {
	uow.registeredNew = append(uow.registeredNew, object)
}

func (uow *UOW) RegisterDirty(object interface{}) {
	uow.registeredDirty = append(uow.registeredDirty, object)
}

func (uow *UOW) Commit() {
	fmt.Println("GOT TO COMMIT")
	fmt.Println(uow.registeredNew)
	for _, e := range uow.registeredNew {
		objectName := reflect.TypeOf(e).String()
		fmt.Println(objectName)
		first := strings.Index(objectName, ".") + 1
		objectName = objectName[first:]
		fmt.Println(objectName)
		switch objectName {
		case "User":
			UserTdg{}.Create(e.(classes.User))
		}
	}
	for _, e := range uow.registeredDirty {
		objectName := reflect.TypeOf(e).String()
		fmt.Println(objectName)
		first := strings.Index(objectName, ".") + 1
		objectName = objectName[first:]
		fmt.Println(objectName)
		switch objectName {
		case "User":
			UserTdg{}.Update(e.(classes.User))
		}
	}
}

//func (uow *UOW)
