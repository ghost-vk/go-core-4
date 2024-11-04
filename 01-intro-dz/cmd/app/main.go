package main

import (
	"fmt"

	"github.com/go-core-4/01-intro-dz/pkg/greetings"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(greetings.Hello("Vlad"))
}
