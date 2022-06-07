package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and flag to disable printing
	// the time, source file, ane line number.
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greeting.Hello("theera")
	// If an error was returned, print it to console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
