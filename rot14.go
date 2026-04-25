package piscine



func Rot14(s string) string {
	result := []rune(s)

	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			result[i] = 'A' + (r-'A'+14)%26
		}
	}

	return string(result)
}