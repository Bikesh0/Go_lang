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

func main() {
	points := &point{}
	setPoint(points)

	// We format the values into a slice of runes
	// 'x = 42, y = 21'
	result := []rune{
		'x', ' ', '=', ' ', rune('0' + points.x/10), rune('0' + points.x%10),
		',', ' ', 'y', ' ', '=', ' ', rune('0' + points.y/10), rune('0' + points.y%10),
		'\n',
	}

	// This single loop uses only 1 PrintRune call
	for _, r := range result {
		z01.PrintRune(r)
	}
}
