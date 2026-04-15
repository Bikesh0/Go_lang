package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	result := RecursiveFactorial(nb - 1)
	if result == 0 {
		return 0
	}

	res := nb * result
	if res < 0 {
		return 0
	}

	return res
}
