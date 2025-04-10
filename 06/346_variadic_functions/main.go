package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4, 5)
	letters := []string{"a"}
	letters = append(letters, "b")
	letters = append(letters, "c", "d", "e", "f", "g")

	twoInts(1, 2)

	myFunc(1, "1", "a")

	severalInts(1)
	severalInts(1, 2, 3)

	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()

	mix(1, true, "a", "b")
	mix(2, false, "a", "b", "c", "d")
}

func twoInts(first int, second int) {
	fmt.Println(first, second)
}

func myFunc(param1 int, param2 ...string) {
	fmt.Println(param1)
	for _, elem := range param2 {
		fmt.Println(elem)
	}
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
