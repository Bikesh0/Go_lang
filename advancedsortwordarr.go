package piscine

// AdvancedSortWordArr sorts a slice of strings using the comparison function f
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)

	// Bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Use the provided function to compare
			if f(a[j], a[j+1]) > 0 {
				// Swap if elements are in the wrong order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
