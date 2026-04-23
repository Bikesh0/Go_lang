package main

import (
	"os"
)

func main() {
	// Check correct number of arguments
	if len(os.Args) != 4 {
		return
	}

	// Validate that inputs are numeric
	if !IsNumeric(os.Args[1]) || !IsNumeric(os.Args[3]) {
		return
	}

	// Convert string arguments to integers
	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])
	op := os.Args[2]

	// Handle division and modulo by zero
	if op == "/" && b == 0 {
		printStr("No division by 0\n")
		return
	}
	if op == "%" && b == 0 {
		printStr("No modulo by 0\n")
		return
	}

	var result int

	// Perform the operation based on the operator
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%":
		result = a % b
	default:
		// Invalid operator → do nothing
		return
	}

	// Print result followed by newline
	PrintNbr(result)
	printStr("\n")
}

// IsNumeric checks if a string represents a valid integer
// Supports optional leading '-' for negative numbers
func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0

	// Handle negative sign
	if s[0] == '-' {
		if len(s) == 1 {
			return false // just "-" is invalid
		}
		i++
	}

	// Check all remaining characters are digits
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// Atoi converts a numeric string to an integer
// Assumes the string is already validated
func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0

	// Handle negative numbers
	if s[0] == '-' {
		sign = -1
		i++
	}

	// Build the integer digit by digit
	for ; i < len(s); i++ {
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}

// PrintNbr prints an integer to standard output
// Uses recursion to print digits in correct order
func PrintNbr(n int) {
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}

	// Print all digits except the last
	if n >= 10 {
		PrintNbr(n / 10)
	}

	// Print last digit
	os.Stdout.Write([]byte{byte(n%10 + '0')})
}

// printStr writes a string to standard output
func printStr(s string) {
	os.Stdout.Write([]byte(s))
}
