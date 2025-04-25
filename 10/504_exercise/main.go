package main

import (
	"fmt"
	"504_exercise/geo"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
