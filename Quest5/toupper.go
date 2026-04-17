package piscine

func ToUpper(str string) string {
	result := ""

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			result += string(str[i] - 32)
		} else {
			result += string(str[i])
		}
	}
	return result
}
