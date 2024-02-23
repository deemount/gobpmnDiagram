package impl

// IFBaseID ...
type IFBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

// IFBaseElement ...
type IFBaseElement interface {
	SetElement(typ string, suffix interface{})
	GetElement() STR_PTR
}

// IFBaseCoordinates ...
type IFBaseCoordinates interface {
	SetCoordinates(x, y int)
	GetCoordinates() (INT_PTR, INT_PTR)
	SetX(x int)
	GetX() INT_PTR
	SetY(y int)
	GetY() INT_PTR
}
