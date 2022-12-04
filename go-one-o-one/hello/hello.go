package main

import (
	"fmt"
	"log"

	"doughdev.com/greetings"
)

func main() {
	// flag to disable printing the time, source file, and line number.
	log.SetFlags(0)
	log.SetPrefix("greetings: ")

	names := []string{
		"Jamie",
		"Maegan",
		"Finn",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
