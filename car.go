package composition

type Car interface {
	Vehicule
	addFuel() string
	park() string
}
type CarImpl struct {
	VehiculeImpl
}

func (CarImpl) addFuel() string {
	return "Adding fuel to car"
}

func (CarImpl) park() string {
	return "Parking car"
}
func (c CarImpl) getMaximumSpeed() int {
	return c.maximumSpeed
}
