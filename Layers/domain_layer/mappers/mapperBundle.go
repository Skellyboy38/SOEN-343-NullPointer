package mappers

type Mappers struct {
	UserMapper *UserMapper
	SessionMapper *SessionMapper
}

var MapperBundle *Mappers

func Init(){
	MapperBundle = &Mappers{InitUserMapper(),InitSessionMapper()}
}
