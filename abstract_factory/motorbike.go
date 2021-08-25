package abstractfactory

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct{}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return 1
}

type CruiseMotorbike struct{}

func (c *CruiseMotorbike) NumWheels() int {
	return 2
}
func (c *CruiseMotorbike) NumSeats() int {
	return 2
}
func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
