package piscine

import (
	"fmt"
)

func QuadB(x, y int) { // x= width , y= length
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Top-left corner, top-right corner, bottom-left corner, bottom-right corner
			if j == 0 && i == 0 {
				fmt.Print("/")
			} else if j == x-1 && i == 0 {
				fmt.Print("\\")
			} else if j == 0 && i == y-1 {
				fmt.Print("\\")
			} else if j == x-1 && i == y-1 {
				fmt.Print("/")

				// Top and bottom edges (excluding corners) and left and right edges (excluding corners)
			} else if (i == 0) || (i == y-1) || (j == x-1 || j == 0) {
				fmt.Print("*")
				// Interior spaces
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
