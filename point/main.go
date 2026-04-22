package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n > 9 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	points := &point{}

	setPoint(points)

	printStr("x = ")
	printNbr(points.x)
	printStr(", y = ")
	printNbr(points.y)
	z01.PrintRune('\n')
}
