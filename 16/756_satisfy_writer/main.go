package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpl, err := template.New("test").Parse("Here's my template!\n")
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}
