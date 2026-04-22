package main

import (
	"os"
)

func printString(s string) {
	for _, c := range s {
		os.Stdout.Write([]byte{byte(c)})
	}
}

func main() {
	args := os.Args[1:]

	// No argument
	if len(args) == 0 {
		printString("File name missing\n")
		return
	}

	// More than one argument
	if len(args) > 1 {
		printString("Too many arguments\n")
		return
	}

	// Exactly one argument
	content, err := os.ReadFile(args[0])
	if err != nil {
		return
	}

	printString(string(content))
}
