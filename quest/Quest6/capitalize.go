package piscine

func Capitalize(s string) string {
	result := []rune(s)
	inWord := false

	for i := 0; i < len(result); i++ {
		r := result[i]

		// Check if alphanumeric
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if !inWord {
				// First character of word → uppercase if letter
				if r >= 'a' && r <= 'z' {
					result[i] = r - 32
				}
				inWord = true
			} else {
				// Inside word → lowercase if letter
				if r >= 'A' && r <= 'Z' {
					result[i] = r + 32
				}
			}
		} else {
			// Not alphanumeric → reset word state
			inWord = false
		}
	}

	return string(result)
}
