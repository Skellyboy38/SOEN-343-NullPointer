package tdg

import "github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"

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

func (uow *UOW)RegisterNew(object interface{}){
	uow.registeredNew = append(uow.registeredNew,object)
}