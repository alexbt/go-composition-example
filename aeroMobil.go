package composition

type AeroMobil struct {
	PlaneImpl
	CarImpl
}

func NewAeroMobil() AeroMobil {
	return AeroMobil{
		PlaneImpl{VehiculeImpl{300, SportEngine{}}, 3000},
		CarImpl{VehiculeImpl{100, StandardEngine{}}}}
}

func (AeroMobil) getModelName() string {
	return "AeroMobil"
}
