package mappers

type Mappers struct {
	UserMapper *UserMapper
}

var MapperBundle *Mappers

func Init(){
	MapperBundle = &Mappers{InitUserMapper()}
}
