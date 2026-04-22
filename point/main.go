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
		string(points.x/'0'+'0') +
		string(points.x%'0'+'0') +
		", y = " +
		string(points.y/'0'+'0') +
		string(points.y%'0'+'0') +
		"\n"

	for _, r := range str {
		z01.PrintRune(r)
	}
}
