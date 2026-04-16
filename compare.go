package piscine

func Compare(a, b string) int {
	i := 0

	for i < len(a) && i < len(b) {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
		i++
	}

	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}