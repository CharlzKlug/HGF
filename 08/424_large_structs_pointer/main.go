package main

import "fmt"

type subscriber struct {
	name string
	rate float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriberD(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscountD(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriberD("Aman Singh")
	applyDiscountD(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSubscriberD("Beth Ryan")
	printInfo(subscriber2)
}
