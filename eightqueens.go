package piscine

import "fmt"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func isSafe(board [8]int, col, row int) bool {
	for i := 0; i < col; i++ {
		// same row
		if board[i] == row {
			return false
		}
		// same diagonal
		if abs(board[i]-row) == col-i {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		fmt.Print(board[i] + 1)
	}
	fmt.Println()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
