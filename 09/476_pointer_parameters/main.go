package main

import "fmt"

type Number int

func (n *Number) Double() {
	*n *= 2
}

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)

	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
}
