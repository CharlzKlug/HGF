package main

import "fmt"

func main() {
	var notes []string
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	fmt.Println(len(notes))
	fmt.Println(len(primes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
