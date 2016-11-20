package mappers

type Mappers struct {
	UserMapper *UserMapper
	ReservationMapper *ReservationMapper
	WaitListMapper *WaitListMapper
}

var MapperBundle *Mappers

func Init(){
	MapperBundle = &Mappers{InitUserMapper(),InitReservationMapper(),InitWaitListMapper()}
}
