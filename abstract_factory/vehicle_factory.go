package abstractfactory

import (
	"fmt"
)

const (
	LuxuryCarType        = 1
	FamilyCarType        = 2
	SportMotorbikeType   = 1
	CruiseMotorbikeType  = 2
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {

	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}

}

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {

	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}
