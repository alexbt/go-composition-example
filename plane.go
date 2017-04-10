package composition

type Plane interface {
	Vehicule
	addFuel() string
	takeoff() string
	land() string
	getMaximumAltitude() int
}

type PlaneImpl struct {
	VehiculeImpl
	maximumAltitude int
}

func (PlaneImpl) addFuel() string {
	return "Adding fuel to plane"
}

func (PlaneImpl) takeoff() string {
	return "PlaneImpl is taking off"
}
func (PlaneImpl) land() string {
	return "PlaneImpl is landing"
}

func (p PlaneImpl) getMaximumAltitude() int {
	return p.maximumAltitude
}

func (p PlaneImpl) getMaximumSpeed() int {
	return p.maximumSpeed
}
