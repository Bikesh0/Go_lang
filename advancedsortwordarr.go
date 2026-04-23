package piscine

// AdvancedSortWordArr sorts the slice using the function f
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)

	// Simple bubble sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Use the comparator function
			if f(a[j], a[j+1]) > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
