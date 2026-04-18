package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // skip program name

	// sort
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[j] < args[i] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	// print WITHOUT filtering
	for _, a := range args {
		for _, ch := range a {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
