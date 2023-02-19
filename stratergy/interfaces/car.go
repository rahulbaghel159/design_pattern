package interfaces

type Car struct {
	OpenBehaviour OpenBehaviour
}

func InitCar(sunroof OpenBehaviour) *Car {
	return &Car{
		OpenBehaviour: sunroof,
	}
}

func (c *Car) SetOpenBehaviour(sunroof OpenBehaviour) {
	c.OpenBehaviour = sunroof
}
