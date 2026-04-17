package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0') // 👉 special case: just print 0
		return
	}

	for n > 0 { // 👉 repeat until all digits are removed

		min := 9 // 👉 start with highest possible digit
		tmp := n // 👉 copy of n (so we don’t destroy original)

		// 🔍 FIND SMALLEST DIGIT
		for tmp > 0 {
			d := tmp % 10 // 👉 get last digit

			if d < min { // 👉 compare with current minimum
				min = d // 👉 update minimum
			}

			tmp /= 10 // 👉 remove last digit (move left)
		}

		z01.PrintRune(rune(min + '0')) // 👉 print smallest digit

		// 🔧 REMOVE ONE OCCURRENCE OF min
		newN := 0        // 👉 rebuilt number
		mult := 1        // 👉 place value (1, 10, 100...)
		removed := false // 👉 ensures only ONE removal

		for n > 0 {
			d := n % 10 // 👉 get last digit

			if d == min && !removed {
				removed = true // 👉 skip this digit ONCE
			} else {
				newN += d * mult // 👉 keep digit
				mult *= 10       // 👉 move to next place
			}

			n /= 10 // 👉 remove last digit
		}

		n = newN // 👉 update n (one digit removed)
	}
}
