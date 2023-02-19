package main

import (
	"fmt"
	"stratergy/interfaces"
)

func main() {
	toy := interfaces.OpenSunroofToyCar{}

	car := interfaces.InitCar(toy)

	fmt.Println("Toy Car")
	car.OpenBehaviour.OpenSunroof()

	car.SetOpenBehaviour(interfaces.OpenSunroofBigCar{})

	fmt.Println("Big Car")
	car.OpenBehaviour.OpenSunroof()
}
