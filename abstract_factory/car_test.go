package abstractfactory

import "testing"

func TestCarFactory(t *testing.T) {

	carF, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	carVevicle, err := carF.Build(LuxuryCarType)

	if err != nil {
		t.Fatal(err)
	}

	luxuryCar, ok := carVevicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}
