package mappers

type Mappers struct {
	UserMapper *UserMapper
	ReservationMapper *ReservationMapper
}

var MapperBundle *Mappers

func Init(){
	MapperBundle = &Mappers{InitUserMapper(),InitReservationMapper()}
}
