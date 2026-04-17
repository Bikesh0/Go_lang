package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for _, a := range arg { // a = string
		for _, ch := range a { // ch = rune
			if (ch >= 'a' && ch <= 'z') ||
				(ch >= 'A' && ch <= 'Z') ||
				(ch >= '0' && ch <= '9') {

				z01.PrintRune(ch)
			}
		}
		z01.PrintRune('\n')
	}
}
