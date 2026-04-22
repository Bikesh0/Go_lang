package main

import (
	"io"
	"os"
)

func printError(err error) {
	os.Stdout.Write([]byte("ERROR: "))
	os.Stdout.Write([]byte(err.Error()))
	os.Stdout.Write([]byte("\n"))
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	// No arguments → stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, name := range args {
		file, err := os.Open(name)
		if err != nil {
			printError(err)
		}

		io.Copy(os.Stdout, file)
		file.Close()
	}
}
