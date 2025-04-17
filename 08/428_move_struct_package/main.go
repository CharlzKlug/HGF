package main

import (
	"428_move_struct_package/magazine"
	"fmt"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
