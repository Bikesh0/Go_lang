package piscine

func Any(f func(string) bool, a []string) bool {
	for _, b := range a {
		if f(b) {
			return true
		}
	}
	return false
}
