package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printRune(r rune) {
	z01.PrintRune(r)
}

func atoi(s string) (int, bool) {
	n := 0
	if s == "" {
		return 0, false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n, ok := atoi(arg)
		if !ok || n < 1 || n > 26 {
			printRune(' ')
			continue
		}

		var letter rune
		if upper {
			letter = rune('A' + n - 1)
		} else {
			letter = rune('a' + n - 1)
		}
		printRune(letter)
	}
	printRune('\n')
}
