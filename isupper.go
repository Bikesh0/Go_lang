package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {

			return false
		}

	}
	return true
}
