// Empty filefunc Atoi(s string) int {
	result := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		// handle sign only at first position
		if i == 0 && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				sign = -1
			}
			continue
		}

		// reject anything not digit
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}