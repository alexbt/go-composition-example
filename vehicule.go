package composition

type Vehicule interface {
	accelerate() string
	getMaximumSpeed() int
	getEngineName() string
}
type VehiculeImpl struct {
	maximumSpeed int
	Engine
}

func (VehiculeImpl) accelerate() string {
	return "VehiculeImpl is accelerating"
}
