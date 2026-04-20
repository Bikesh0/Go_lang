package piscine

// we are using a SortIntegerTable to sort a slice of integers in ascending order which is []int
// new to learn Bubble Sort algorithm.
func SortIntegerTable(table []int) {
	n := len(table)
	// we declare n to be the lenght of table or list that will be given or included already.

	// After each pass, the largest unsorted element moves to its correct position
	for i := 0; i < n-1; i++ { // starting for look Outer loop with index i untill the lenth of string or inputs or table in the slice

		// The "-i-1" avoids re-checking elements already sorted at the end
		for j := 0; j < n-i-1; j++ { // Inner loop compares adjacent elements like first with this second if i is first j is second element

			if table[j] > table[j+1] { // so we compare here and interchange position in ascending order.
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
} // works well when tested
