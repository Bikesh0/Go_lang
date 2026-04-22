package piscine

func Split(s, sep string) []string {
	var result []string

	if sep == "" {
		// If separator is empty, return the whole string as one element
		return []string{s}
	}

	start := 0
	for i := 0; i <= len(s)-len(sep); {
		if s[i:i+len(sep)] == sep {
			// Add substring before separator
			result = append(result, s[start:i])
			// Move past separator
			i += len(sep)
			start = i
		} else {
			i++
		}
	}

	// Add remaining part
	result = append(result, s[start:])

	return result
}
