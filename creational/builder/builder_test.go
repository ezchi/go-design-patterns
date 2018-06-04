package builder

import "testing"

func TestBuilderPatter(t *testing.T) {
	manufactoringComplex := ManufactoringDirector{}

	carBuilder := &CarBuilder{}

	manufactoringComplex.SetBuilder(carBuilder)
	manufactoringComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("wheels on a car must be 4, but they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuiler := &BikeBuilder{}

	manufactoringComplex.SetBuilder(bikeBuiler)
	manufactoringComplex.Construct()

	motorbike := bikeBuiler.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("wheels on motobike must be 2 and they where %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}
}
