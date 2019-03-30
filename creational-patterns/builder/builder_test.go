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

	if motorbike.Wheels != 2 {
		t.Errorf("wheels on motobike must be 2 and they where %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}

	if motorbike.Seats != 2 {
		t.Errorf("seats on a motorbike must be 2 and they were %d\n", motorbike.Seats)
	}

	busBuilder := &BusBuilder{}

	manufactoringComplex.SetBuilder(busBuilder)
	manufactoringComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("wheels on bus must be 8 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("structure on a bus must be 'Bus' and was %s\n", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("seats on a bus must be 30 and they were %d\n", bus.Seats)
	}
}
