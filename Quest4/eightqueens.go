package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	var backtrack func(int)
	backtrack = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '1'))
			}
			z01.PrintRune('\n')
			return
		}

		for row := 0; row < 8; row++ {
			valid := true

			for i := 0; i < col; i++ {
				diff := board[i] - row
				if diff < 0 {
					diff = -diff
				}
				if board[i] == row || diff == col-i {
					valid = false
					break
				}
			}

			if valid {
				board[col] = row
				backtrack(col + 1)
			}
		}
	}

	backtrack(0)
}
