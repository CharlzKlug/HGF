package main

import "fmt"

type subscriber struct {
	rate float32
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
}
