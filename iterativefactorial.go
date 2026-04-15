package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	factorial := 1
	for i := 1; i <= nb; i++ {
		factorial *= i
		if factorial < 0 {
			return 0
		}
	}
	return factorial
}
