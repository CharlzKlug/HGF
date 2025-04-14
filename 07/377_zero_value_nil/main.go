package main

import "fmt"

func main() {
	var myMap map[int]string = make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)

	var nilMap map[int]string
	fmt.Printf("%#v\n", nilMap)
	nilMap[3] = "three"
}
