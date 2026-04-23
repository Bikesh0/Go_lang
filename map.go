package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, b := range a {
		result[i] = f(b)
	}
	return result
}
