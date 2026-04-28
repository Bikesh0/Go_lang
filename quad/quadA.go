package piscine

import (
	"fmt"
)

func QuadA(x, y int) { // x= width , y= length
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ { // i = no. of row and j = no. ofcolumn
		for j := 0; j < x; j++ {
			if (j == 0 && i == 0) || (j == x-1 && i == 0) || (j == 0 && i == y-1) || (j == x-1 && i == y-1) { // corners
				fmt.Print("o")
			} else if j == x-1 || j == 0 { //
				fmt.Print("|")
			} else if (i == 0) || (i == y-1) { // top and bottom edges (excluding corners)
				fmt.Print("-")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
