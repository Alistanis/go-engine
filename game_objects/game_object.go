package game_object


type GameObject struct {
	
}

type Object interface {
	ChangeState(obj *GameObject)
}



