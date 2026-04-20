package piscine

// ConvertBase converts nbr from baseFrom to baseTo.
func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := toDecimal(nbr, baseFrom)
	return toBase(decimal, baseTo)
}

// -------------------------
// Convert from baseFrom → decimal
// -------------------------
func toDecimal(nbr, base string) int {
	baseLen := len(base)
	result := 0

	for i := 0; i < len(nbr); i++ {
		value := index(base, nbr[i])
		result = result*baseLen + value
	}

	return result
}

// -------------------------
// Convert from decimal → baseTo
// -------------------------
func toBase(nbr int, base string) string {
	if nbr == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	var res []byte

	for nbr > 0 {
		remainder := nbr % baseLen
		res = append([]byte{base[remainder]}, res...)
		nbr /= baseLen
	}

	return string(res)
}

// -------------------------
// Find index of a character in a base
// -------------------------
func index(base string, c byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return 0
}
