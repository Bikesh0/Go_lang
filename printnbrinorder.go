package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Loop through digits from smallest to largest
	for i := 0; i <= 9; i++ {

		tmp := n // copy so we don’t destroy n

		// scan all digits of n
		for tmp > 0 {
			d := tmp % 10

			if d == i {
				z01.PrintRune(rune(d + '0'))
			}

			tmp /= 10
		}
	}
}
