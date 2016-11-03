package tdg

import "github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/classes"

type user classes.User

type TDGqueue []TDG

type UOW struct {
	registeredNew TDGqueue
	registeredDirty TDGqueue
	registeredDeleted TDGqueue
}

//func RegisterNew(object TDG)