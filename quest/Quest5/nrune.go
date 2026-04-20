package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	for i, r := range s {
		if i == n-1 {
			return r
		}
	}
	return 0
}
