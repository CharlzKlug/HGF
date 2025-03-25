package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return (area / 10.0)
}

func main() {
	fmt.Println(paintNeeded(3, 2))
}
