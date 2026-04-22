package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 'Z' - '0'                 // 42
	ptr.y = ('A' - '0') + ('K' - 'A') // 21
}

func main() {
	points := &point{}
	setPoint(points)

	str := "x = " +
		string('0'+(points.x/10)) +
		string('0'+(points.x%10)) +
		", y = " +
		string('0'+(points.y/10)) +
		string('0'+(points.y%10)) +
		"\n"

	for _, r := range str {
		z01.PrintRune(r)
	}
}
