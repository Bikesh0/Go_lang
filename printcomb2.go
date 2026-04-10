package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {

			print2(a)
			z01.PrintRune(' ')
			print2(b)

			if !(a == 98 && b == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func print2(n int) {
	printDigit(n / 10)
	printDigit(n % 10)
}

func printDigit(d int) {
	if d == 0 {
		z01.PrintRune('0')
	}
	if d == 1 {
		z01.PrintRune('1')
	}
	if d == 2 {
		z01.PrintRune('2')
	}
	if d == 3 {
		z01.PrintRune('3')
	}
	if d == 4 {
		z01.PrintRune('4')
	}
	if d == 5 {
		z01.PrintRune('5')
	}
	if d == 6 {
		z01.PrintRune('6')
	}
	if d == 7 {
		z01.PrintRune('7')
	}
	if d == 8 {
		z01.PrintRune('8')
	}
	if d == 9 {
		z01.PrintRune('9')
	}
}
