package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	pointer.myField = 9
	fmt.Println(pointer.myField)
}
