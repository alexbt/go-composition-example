package composition

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCar(t *testing.T) {
	//f := Ferrari{CarImpl{150}}
	f := NewFerrari()
	assert.Equal(t, "Ferrari", f.getModelName());
	assert.Equal(t, "Adding fuel to car", f.addFuel());
	assert.Equal(t, 150, f.getMaximumSpeed());
	assert.Equal(t, "Parking car", f.park());
	assert.Equal(t, "VehiculeImpl is accelerating", f.accelerate());
	assert.Equal(t, "SportEngine", f.getEngineName());
}

func TestBoeing(t *testing.T) {
	//b := Boeing{PlaneImpl{500}}
	b := NewBoeing()
	assert.Equal(t, "Boeing", b.getModelName());
	assert.Equal(t, "Adding fuel to plane", b.addFuel());
	assert.Equal(t, 500, b.getMaximumSpeed());
	assert.Equal(t, 10000, b.getMaximumAltitude());
	assert.Equal(t, "PlaneImpl is landing", b.land());
	assert.Equal(t, "PlaneImpl is taking off", b.takeoff());
	assert.Equal(t, "VehiculeImpl is accelerating", b.accelerate());
	assert.Equal(t, "SportEngine", b.getEngineName());
}

func TestAeroMobil(t *testing.T) {
	//a := AeroMobil{PlaneImpl{300}, CarImpl{100}}
	a := NewAeroMobil()
	assert.Equal(t, "AeroMobil", a.getModelName());
	assert.Equal(t, "Adding fuel to car", a.CarImpl.addFuel());
	assert.Equal(t, "Adding fuel to plane", a.PlaneImpl.addFuel());
	assert.Equal(t, 100, a.CarImpl.getMaximumSpeed());
	assert.Equal(t, 300, a.PlaneImpl.getMaximumSpeed());
	assert.Equal(t, 3000, a.getMaximumAltitude());
	assert.Equal(t, "PlaneImpl is landing", a.land());
	assert.Equal(t, "PlaneImpl is taking off", a.takeoff());
	assert.Equal(t, "Parking car", a.park());
	assert.Equal(t, "VehiculeImpl is accelerating", a.CarImpl.accelerate());
	assert.Equal(t, "VehiculeImpl is accelerating", a.PlaneImpl.accelerate());
	assert.Equal(t, "SportEngine", a.PlaneImpl.getEngineName());
	assert.Equal(t, "StandardEngine", a.CarImpl.getEngineName());
}
