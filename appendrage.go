package piscine

func AppendRange(min, max int) []int {
	// If min is greater than or equal to max,
	// there are no valid numbers to include in the slice.
	// According to the requirement, we return a nil slice.
	if min >= max {
		return nil
	}

	// Declare a slice of integers.
	// At this point, 'result' is a nil slice (empty, no memory allocated yet).
	var result []int

	// Loop from min up to (but NOT including) max.
	// Example: if min = 5 and max = 10 → i will be 5,6,7,8,9
	for i := min; i < max; i++ {

		// Append the current value of i to the slice.
		// If the slice is nil, append will automatically
		// allocate memory and add the first element.
		result = append(result, i)
	}

	// Return the final slice containing all values from min to max-1.
	return result
}
