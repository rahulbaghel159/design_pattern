package interfaces

import "fmt"

type OpenSunroofBigCar struct{}

func (bc OpenSunroofBigCar) OpenSunroof() {
	fmt.Println("Opening in style")
}
