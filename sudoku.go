package main

import (
	"fmt"
	"os"
)

const size = 9 // Standard Sudoku board size (9x9)

func main() {
	// Expect exactly 9 rows as arguments (plus program name)
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	// Create a 2D slice to represent the Sudoku board
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		// Each row must have exactly 9 characters
		if len(os.Args[i+1]) != 9 {
			fmt.Println("Error")
			return
		}
		//  part 2 Allocate row memory
		board[i] = make([]int, size)
		// Loop in each column
		for j := 0; j < size; j++ {
			// Position of character
			character := os.Args[i+1][j]

			// '.' represents an empty cell
			if character == '.' {
				board[i][j] = 0
			} else if character >= '1' && character <= '9' {
				// Convert character digit to integer
				board[i][j] = int(character - '0')
			} else {
				// Invalid character found
				fmt.Println("Error")
				return
			}
		}
	}

	// Part 3 Check if the initial board configuration is valid
	if !isValidBoard(board) {
		fmt.Println("Error")
		return
	}

	// Count number of solutions
	count := 0
	solveCount(board, &count)

	// Must have exactly one solution
	if count != 1 {
		fmt.Println("Error")
		return
	}

	// Solve again to fill the board with the solution
	if !solve(board) {
		fmt.Println("Error")
		return
	}

	// Print the solved board
	printBoard(board)
}

// Solves the Sudoku using backtracking
func solve(board [][]int) bool {
	row, col, found := findEmpty(board)

	// If no empty cell is found, the board is solved
	if !found {
		return true
	}

	// Try numbers 1 through 9
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) { //continue if no repeating number
			board[row][col] = num //temp fill the cell

			// Recursively attempt to solve the rest
			if solve(board) {
				return true
			}

			// Backtrack if solution fails
			board[row][col] = 0
		}
	}

	return false
}

// Counts the number of possible solutions using backtracking
func solveCount(board [][]int, count *int) {
	// Stop early if more than one solution is found
	if *count > 1 {
		return
	}

	row, col, found := findEmpty(board)

	// If no empty cell is found, a solution is complete
	if !found {
		*count++
		return
	}

	// Try numbers 1 through 9
	for num := 1; num <= 9; num++ {
		if isSafe(board, row, col, num) { // no duplicate or repeating number
			board[row][col] = num //temp value

			// Recursively count solutions
			solveCount(board, count) //recurseive moves forward to solve the board

			// Backtrack
			board[row][col] = 0 // reset either solving or not solving. addcount if soln found
		}
	}
}

// Finds the next empty cell (value 0)
func findEmpty(board [][]int) (int, int, bool) {
	for i := 0; i < size; i++ { // loop through row
		for j := 0; j < size; j++ { // loop through column
			if board[i][j] == 0 { // check for empty cell
				return i, j, true // found empty cell at postion i,i. stops searching other empty cell
			}
		}
	}
	return 0, 0, false // no empty cell found in entire board
}

// Checks if the initial board is valid
func isValidBoard(board [][]int) bool {
	for i := 0; i < size; i++ { // lopping throught rows
		for j := 0; j < size; j++ { // looping thorugh columns
			if board[i][j] != 0 {
				value := board[i][j] // save current number
				board[i][j] = 0      // Temporarily set value 0

				// Check if placing val back is safe
				if !isSafe(board, i, j, value) {
					board[i][j] = value // restore before returning
					return false
				}

				board[i][j] = value // Restore original value
			}
		}
	}
	return true
}

// Checks if placing 'num' at (row, col) is valid in small grid
func isSafe(board [][]int, row, col, num int) bool {
	// Check the row
	for i := 0; i < size; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check the column
	for i := 0; i < size; i++ { // check in entire row if there duplicate numbers
		if board[i][col] == num { // num exoist false
			return false
		}
	}

	// Find top-left corner of the 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3

	// Check the 3x3 box (simple nested loops)
	for i := 0; i < 3; i++ { // check 9 cells for num exit alresy
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

// Prints the Sudoku board in a readable format
func printBoard(board [][]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(board[i][j]) // print num at posit of cell
			if j != size-1 {       // Print space until j is not last num
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
