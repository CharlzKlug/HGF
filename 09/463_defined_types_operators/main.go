package main

import "fmt"

type Liters float64
type Gallons float64
type Title string

func main() {
	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	fmt.Println(Title("Alien") + "s")
	// fmt.Println(Title("Jaws 2") - " 2")

	// fmt.Println(Liters(1.2) + Gallons(3.4))
	// fmt.Println(Gallons(1.2) == Liters(1.2))
}
