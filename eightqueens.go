package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var game [8]int

	var backtrack func(int)
	backtrack = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(game[i] + '1'))
			}
			z01.PrintRune('\n')
			return
		}

		for row := 0; row < 8; row++ {
			valid := true

			for i := 0; i < col; i++ {
				if game[i] == row || game[i]-i == row-col || game[i]+i == row+col {
					valid = false
					break
				}
			}

			if valid {
				game[col] = row
				backtrack(col + 1)
			}
		}
	}

	backtrack(0)
}
