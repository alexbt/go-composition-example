package composition

type Bmw struct {
	CarImpl
}

func (Bmw) getModelName() string {
	return "BMW"
}

func NewBmw() Bmw {
	return Bmw{CarImpl{VehiculeImpl{150, SportEngine{}}}}
}
