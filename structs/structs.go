package structs

import (
	"fmt"
	"go_tutorials/engines"
)

func Structs() {
	var myEngine engines.GasEngine = engines.GasEngine{Mpg: 25, Gallons: 15}
	engines.CanMakeIt(myEngine, 50)
	fmt.Printf("Total miles left in tank: %v", myEngine.MilesLeft())
	fmt.Println(myEngine.Mpg, myEngine.Gallons)
	var myEngine2 = struct {
		Mpg     uint8
		Gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.Mpg, myEngine2.Gallons)
}
