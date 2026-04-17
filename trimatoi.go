package piscine

func TrimAtoi(s string) int {
	result := 0         // 👉 final number
	sign := 1           // 👉 default is positive
	foundDigit := false // 👉 track if we found any digit yet

	for _, ch := range s {

		// 🔹 Check for minus sign BEFORE digits
		if ch == '-' && !foundDigit {
			sign = -1
		}

		// 🔹 Check if character is a digit
		if ch >= '0' && ch <= '9' {

			foundDigit = true // 👉 we now started reading digits

			// Convert rune → int
			digit := int(ch - '0')

			// Build number
			result = result*10 + digit
		}
	}

	// 🔹 If no digits found → return 0
	if !foundDigit {
		return 0
	}

	return result * sign
}
