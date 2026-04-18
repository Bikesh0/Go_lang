package piscine

func SplitWhiteSpaces(s string) []string {
	// Slice to store the resulting words
	var result []string

	// Temporary string to build each word character by character
	word := ""

	// Loop through each character in the string
	for i := 0; i < len(s); i++ {
		// Check if current character is NOT a separator
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			// Add character to current word
			word += string(s[i])
		} else {
			// If we hit a separator and word is not empty
			if word != "" {
				// Add completed word to result slice
				result = append(result, word)
				// Reset word for next use
				word = ""
			}
		}
	}

	// After the loop, check if there's a last word (no separator at end)
	if word != "" {
		result = append(result, word)
	}

	// Return slice of words
	return result
}
