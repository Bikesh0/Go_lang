package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := Atoi(os.Args[1])
	b, ok2 := Atoi(os.Args[3])
	op := os.Args[2]

	if !ok1 || !ok2 {
		return
	}

	if op == "/" && b == 0 {
		printStr("No division by 0\n")
		return
	}
	if op == "%" && b == 0 {
		printStr("No modulo by 0\n")
		return
	}

	var res int64

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "%":
		res = a % b
	default:
		return
	}

	// check overflow for int
	maxInt := int64(int(^uint(0) >> 1))
	minInt := -maxInt - 1

	if res > maxInt || res < minInt {
		return
	}

	PrintNbr(int(res))
	printStr("\n")
}

// Atoi returns int64 directly
func Atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	i := 0
	var res int64

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
	}

	return res * sign, true
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
