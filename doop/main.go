package main

import (
	"os"
)

const (
	maxInt = 9223372036854775807
	minInt = -9223372036854775808
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := Atoi(os.Args[1])
	b, ok2 := Atoi(os.Args[3])
	op := os.Args[2]

	// invalid numbers
	if !ok1 || !ok2 {
		return
	}

	// division / modulo by zero
	if op == "/" && b == 0 {
		printStr("No division by 0\n")
		return
	}
	if op == "%" && b == 0 {
		printStr("No modulo by 0\n")
		return
	}

	var result int
	var ok bool = true

	switch op {
	case "+":
		if (b > 0 && a > maxInt-b) || (b < 0 && a < minInt-b) {
			ok = false
		} else {
			result = a + b
		}
	case "-":
		if (b < 0 && a > maxInt+b) || (b > 0 && a < minInt+b) {
			ok = false
		} else {
			result = a - b
		}
	case "*":
		if a != 0 && (b > maxInt/a || b < minInt/a) {
			ok = false
		} else {
			result = a * b
		}
	case "/":
		result = a / b
	case "%":
		result = a % b
	default:
		return
	}

	if !ok {
		return
	}

	PrintNbr(result)
	printStr("\n")
}

func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0
	var res int64 = 0

	if s[0] == '-' {
		sign = -1
		i++
		if len(s) == 1 {
			return 0, false
		}
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		res = res*10 + int64(s[i]-'0')

		// overflow check during parsing
		if sign == 1 && res > maxInt {
			return 0, false
		}
		if sign == -1 && -res < minInt {
			return 0, false
		}
	}

	return int(res) * sign, true
}

func PrintNbr(n int) {
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	os.Stdout.Write([]byte{byte(n%10 + '0')})
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}
