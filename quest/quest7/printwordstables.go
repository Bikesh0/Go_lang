package piscine

import "github.com/01-edu/z01"

// PrintWordsTables prints each string of the slice
// on a separate line.
func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		// Print each character of the word
		for _, r := range a[i] {
			z01.PrintRune(r)
		}
		// New line after each word
		z01.PrintRune('\n')
	}
}
