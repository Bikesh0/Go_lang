package piscine

func Max(a []int) int {
	highestValue := 0

	for i := 0; i < len(a); i++ {
		if a[i] > highestValue {
			highestValue = a[i]
		}
	}
	return highestValue
}
