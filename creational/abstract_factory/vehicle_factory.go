package abstract_factory

import "fmt"

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCarType), nil

	case FamilyCarType:
		return new(FamilyCarType), nil

	default:
		return nil, fmt.Errorf("vehicle of type %d not reconized\n", v)
	}
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbikebike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil

	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}
