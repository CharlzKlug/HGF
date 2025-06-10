package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)
	greeting(myChannel)
	// fmt.Println(<-myChannel)
}
