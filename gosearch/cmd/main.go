package main

import (
	"flag"
	"fmt"
)

func main() {
	in := parseInput()
	fmt.Println(*in.s)
}

func parseInput() Input {
	search := flag.String("s", "", "search word")
	flag.Parse()

	return Input{
		s: search,
	}
}

type Input struct {
	s *string
}
