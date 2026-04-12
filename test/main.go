package main

import "fmt"

// we are using a SortIntegerTable to sort a slice of integers in ascending order which is []int
// new to learn Bubble Sort algorithm.
func SortIntegerTable(table []int) {
	n := len(table)
	// we declare n to be the length of table

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

// ✅ THIS is what you're missing
func main() {
	arr := []int{5, 3, 1, 4, 2}

	SortIntegerTable(arr)

	fmt.Println(arr)
}
