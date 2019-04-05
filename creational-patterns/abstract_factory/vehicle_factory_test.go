package abstract_factory

import "testing"

func TestMotorbike(t *testing.T) {
	factory, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	sportMotorbike, err := factory.NewVehicle(SportMotorbikeType)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Sport motorbike has %d wheels and %d seats\n", sportMotorbike.NumWheels(), sportMotorbike.NumSeats())

	sportBike, ok := sportMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	factory, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	car, err := factory.NewVehicle(LuxuryCarType)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", car.NumWheels())

	luxuryCar, ok := car.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors.\n", luxuryCar.NumDoors())
}
