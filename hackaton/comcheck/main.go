package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, r := range os.Args[1:] {
			if r == "01" || r == "galaxy" || r == "galaxy 01" {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
