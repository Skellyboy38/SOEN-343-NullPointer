package tdg

import (
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"
	"reflect"
	"fmt"
)

type user classes.User

type objectQueue []interface{}

var UOWSingleTon UOW

type UOW struct {
	registeredNew objectQueue
	registeredDirty objectQueue
	registeredDeleted objectQueue
}

func InitUOW(){
	UOWSingleTon = UOW{objectQueue{},objectQueue{},objectQueue{}}
}

func (uow *UOW) RegisterNew(object interface{}){
	uow.registeredNew = append(uow.registeredNew,object)
}

func (uow *UOW) Commit (){
	for _, e := range uow.registeredNew{
		fmt.Print(reflect.ValueOf(e))
		//switch reflect.ValueOf(e){
		//case
		//}
	}
}

//func (uow *UOW)