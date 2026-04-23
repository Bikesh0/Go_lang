package main

import (
	"fmt"
)

// Compare compares two strings using ASCII order
func Compare(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// AdvancedSortWordArr sorts using the comparison function f
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)

	// Bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Use f to decide order
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}

	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}
