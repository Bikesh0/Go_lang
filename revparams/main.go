package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := len(args) - 1; i > len(args); i-- {
		for _, a := range args[i] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
