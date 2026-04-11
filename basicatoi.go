package piscine

func BasicAtoi(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return result
}
