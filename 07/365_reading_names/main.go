// count tallies the number of times each line
// occurs within a file.
package main

import (
	"fmt"
	"365_reading_names/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
