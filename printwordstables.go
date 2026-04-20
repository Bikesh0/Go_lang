package piscine

// SplitWhiteSpaces splits a string into words separated by
// spaces, tabs, or newlines and returns them as a slice.
func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string

	for i := 0; i < len(s); i++ {
		c := s[i]

		// If we hit a separator, save the current word (if any)
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			// Build current word
			word += string(c)
		}
	}

	// Add last word if string didn't end with a separator
	if word != "" {
		result = append(result, word)
	}

	return result
}
