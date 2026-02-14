package engines

import "fmt"

type GasEngine struct {
	Mpg     uint8
	Gallons uint8
}

type ElectricEngine struct {
	Mpkwh uint8
	Kwh   uint8
}

func (e ElectricEngine) MilesLeft() uint8 {
	return e.Mpkwh * e.Kwh
}

func (e GasEngine) MilesLeft() uint8 {
	return e.Gallons * e.Mpg
}

type Engine interface {
	MilesLeft() uint8
}

func CanMakeIt(e Engine, miles uint8) {
	if miles <= e.MilesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}
