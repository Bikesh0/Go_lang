package piscine

// SortWordArr sorts a slice of strings in ascending ASCII order.
func SortWordArr(a []string) {
	n := len(a)

	// Bubble sort algorithm
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare strings directly (ASCII order)
			if a[j] > a[j+1] {
				// Swap elements
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
