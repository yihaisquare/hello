package main

import (
	"fmt"
	"log"

	"github.com/yihaisquare/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	fmt.Println("hello world")

	fmt.Println(quote.Go())
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
