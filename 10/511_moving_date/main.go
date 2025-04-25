package main

import (
	"fmt"
	"511_moving_date/calendar"
)

func main() {
	date := calendar.Date{}
	date.year = 2019
	date.month = 14
	date.day = 50
	fmt.Println(date)

	date = calendar.Date{year: 0, month: 0, day: -2}
	fmt.Println(date)
}
