package main

import (
	"511_moving_date/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	event := calendar.Event{}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

	fmt.Println(event.Date.Year())
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())

	eventBirthday := calendar.Event{}
	eventBirthday.SetTitle("Mom's birthday djkfksjdfkjdkfjdkjfskjfkdjskfjksdjfksdjfkjsdkjfksdjfkjsdkf")
	fmt.Println(eventBirthday.Title())
}
