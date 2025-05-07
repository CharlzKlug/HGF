package main

import "fmt"

type Appliance interface {
	TurnOn() error
}

type Fan string

func (f Fan) TurnOn() error {
	fmt.Println("Spinning")
	return nil
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	// device = CoffeePot("LuxBrew")
	device.TurnOn()
}
