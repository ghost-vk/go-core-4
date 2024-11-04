package main

import (
	"fmt"
	"log"

	"github.com/go-core-4/01-intro-dz/pkg/greetings"
	"rsc.io/quote/v4"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("01-intro-dz: ")
	log.SetFlags(0)

	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println("\n")

	// Request a greeting message.
	_, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Println(err)
	}

	helloMsg, _ := greetings.Hello("Vlad")
	fmt.Println(helloMsg)
}
