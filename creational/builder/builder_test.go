package builder

import "testing"

func TestBuildPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := new(CarBuilder)
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Constructor()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := new(BikeBuilder)

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Constructor()

	motorBike := bikeBuilder.GetVehicle()

	motorBike.Seats = 1

	if motorBike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorBike.Wheels)
	}

	if motorBike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorBike.Structure)
	}
}
