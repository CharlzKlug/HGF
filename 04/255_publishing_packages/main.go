package main

import (
	"255_publishing_packages/github.com/headfirstgo/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Grade is", grade)
}
