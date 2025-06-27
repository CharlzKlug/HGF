package main

import "fmt"

func sayHi(name string) {
	fmt.Println("Hi", name)
}

func main() {
	var myFunction func(string)
	myFunction = sayHi
	myFunction("George")
}
