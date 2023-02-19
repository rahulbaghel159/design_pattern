package interfaces

import "fmt"

type OpenSunroofToyCar struct{}

func (tc OpenSunroofToyCar) OpenSunroof() {
	fmt.Println("Not possible")
}
