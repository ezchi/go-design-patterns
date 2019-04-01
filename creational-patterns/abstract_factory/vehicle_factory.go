package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleType int

const (
	LuxuryCarType VehicleType = iota
	FamilyCarType
	SportMotorbikeType
	CruiseMotorbikeType
)

type VehicleFactory interface {
	NewVehicle(v VehicleType) (Vehicle, error)
}

type FactoryType int

const (
	CarFactoryType FactoryType = iota
	MotorbikeFactoryType
)

func BuildFactory(f FactoryType) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil

	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil

	default:
		return nil, errors.New(fmt.Sprintf("Factory with type %d is not reconized\n", f))
	}
}

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v VehicleType) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil

	case FamilyCarType:
		return new(FamilyCar), nil

	default:
		return nil, fmt.Errorf("vehicle of type %d not reconized\n", v)
	}
}

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v VehicleType) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil

	case CruiseMotorbikeType:
		// return new(CruiseMotorbike), nil
		return nil, errors.New("CruiseMotorbike is not implemented")

	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized\n", v)
	}
}
