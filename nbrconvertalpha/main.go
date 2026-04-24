package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		n, ok := atoi(arg)
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}

		if upper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
	}
	z01.PrintRune('\n')
}
