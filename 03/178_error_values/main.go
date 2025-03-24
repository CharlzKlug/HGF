package main

import (
	// "errors"
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.33333)
	fmt.Println(err)
	log.Fatal(err)
}
