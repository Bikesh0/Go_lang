package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		j := 0

		for j < len(toFind) && s[i+j] == toFind[j] {
			j++
		}

		if j == len(toFind) {
			return i
		}
	}

	return -1
}
