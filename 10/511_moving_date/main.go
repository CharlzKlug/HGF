package main

import (
	"fmt"
	"511_moving_date/calendar"
)

func main() {
	date := calendar.Date{}
	date.Year = 2019
	date.Month = 14
	date.Day = 50
	fmt.Println(date)

	date = calendar.Date{Year: 0, Month: 0, Day: -2}
	fmt.Println(date)
}
