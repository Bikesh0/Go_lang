package piscine

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, v := range a {
		counts[v]++
	}

	for k, v := range counts {
		if v%2 != 0 {
			return k
		}
	}

	return -1
}
