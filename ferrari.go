package composition

type Ferrari struct {
	CarImpl
}

func (Ferrari) getModelName() string {
	return "Ferrari"
}

func NewFerrari() Ferrari {
	return Ferrari{CarImpl{VehiculeImpl{150, SportEngine{}}}}
}
