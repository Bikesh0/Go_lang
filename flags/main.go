package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortASCII(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return string(arr)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	str := ""

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insert = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			str = arg
		}
	}

	// Apply insert first
	if insert != "" {
		str += insert
	}

	// Then sort
	if order {
		str = sortASCII(str)
	}

	fmt.Println(str)
}
