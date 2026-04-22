package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printError(err error, name string) {
	fmt.Println("ERROR:", err)
	os.Exit(1)
}

func readAndPrint(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	args := os.Args[1:]

	// No arguments → read stdin
	if len(args) == 0 {
		readAndPrint(os.Stdin)
		return
	}

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printError(err, name)
		}

		readAndPrint(file)
		file.Close()
	}
}
