package piscine

// SortIntegerTable sorts a slice of integers in ascending order
// using the Bubble Sort algorithm.
func SortIntegerTable(table []int) {
	n := len(table)

	// Outer loop controls how many passes we make over the slice
	// After each pass, the largest unsorted element moves to its correct position
	for i := 0; i < n-1; i++ {

		// Inner loop compares adjacent elements
		// The "-i-1" avoids re-checking elements already sorted at the end
		for j := 0; j < n-i-1; j++ {

			// If the current element is greater than the next one, swap them
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
