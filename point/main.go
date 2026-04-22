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

	z := '0'
	a := z
	a++ // 1
	a++ // 2

	b := a
	b++ // 3
	b++ // 4

	output := []rune{
		'x', ' ', '=', ' ',
		b, a, // 4 2
		',', ' ',
		'y', ' ', '=', ' ',
		a, z + 1, // 2 1
		'\n',
	}

	for _, c := range output {
		z01.PrintRune(c)
	}
}
