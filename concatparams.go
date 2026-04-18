package piscine

func ConcatParams(args []string) string {
	// Declare a string variable to store the final result
	var result string

	// Loop through all elements in the slice 'args'
	for i := 0; i < len(args); i++ {

		// Add the current string to the result
		result += args[i]

		// Add a newline ONLY if it's NOT the last element
		// This avoids having an extra newline at the end
		if i != len(args)-1 {
			result += "\n"
		}
	}

	// Return the final concatenated string
	return result
}
