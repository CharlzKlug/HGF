package main

import (
	"428_move_struct_package/magazine"
	"fmt"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
	subscribr := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscribr.Name)
	fmt.Println("Rate:", subscribr.Rate)
	fmt.Println("Active:", subscribr.Active)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
