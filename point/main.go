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
	p := &point{}
	setPoint(p)

	// print "x = "
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// print x (42)
	z01.PrintRune(p.x/10 + '0')
	z01.PrintRune(p.x%10 + '0')

	// print ", y = "
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	// print y (21)
	z01.PrintRune(p.y/10 + '0')
	z01.PrintRune(p.y%10 + '0')

	z01.PrintRune('\n')
}
