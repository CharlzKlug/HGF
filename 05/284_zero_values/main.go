package main
import "fmt"
func main() {
	var primes [5]int
	primes[0] = 2
	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])
	// fmt.Println(primes[5])

	var notes [7]string
	notes[0] = "do"
	fmt.Println(notes[3])
	fmt.Println(notes[6])
	fmt.Println(notes[0])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	var bars [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(bars[3], bars[6], bars[0])
	var somePrimes [5] int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(somePrimes[0], somePrimes[2], somePrimes[4])

	text := [3]string {"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}
	fmt.Println(text[0])
}
