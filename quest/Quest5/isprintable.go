package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			return false
		}
	}
	return true
}
