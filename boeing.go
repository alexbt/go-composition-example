package composition

type Boeing struct {
	PlaneImpl
}

func (Boeing) getModelName() string {
	return "Boeing"
}

func NewBoeing() Boeing {
	return Boeing{PlaneImpl{VehiculeImpl{500, SportEngine{}}, 10000}, }
}
