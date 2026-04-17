package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]

	sort.Strings(args)

	for _, s := range args {
		fmt.Println(s)
	}
}
