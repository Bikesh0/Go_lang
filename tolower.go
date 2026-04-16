package piscine

func ToLower(str string) string {
	result := ""

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			result += string(str[i] + 32)
		} else {
			result += string(str[i])
		}
	}
	return result
}
