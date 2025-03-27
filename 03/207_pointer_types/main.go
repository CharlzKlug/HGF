package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// var myInt int
	// fmt.Println(reflect.TypeOf(&myInt))
	// var myFloat float64
	// fmt.Println(reflect.TypeOf(&myFloat))
	// var myBool bool
	// fmt.Println(reflect.TypeOf(&myBool))

	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
}
