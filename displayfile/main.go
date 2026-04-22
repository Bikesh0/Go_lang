package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		return
	}

	n, err := strconv.Atoi(args[1])
	if err != nil || n < 0 {
		return
	}

	files := args[2:]
	exitCode := 0

	for i, name := range files {
		file, err := os.Open(name)
		if err != nil {
			printString(err.Error() + "\n")
			exitCode = 1
			continue
		}

		info, err := file.Stat()
		if err != nil {
			file.Close()
			exitCode = 1
			continue
		}

		size := info.Size()
		start := int64(0)

		if int64(n) < size {
			start = size - int64(n)
		}

		file.Seek(start, 0)

		// Header if multiple files
		if len(files) > 1 {
			if i > 0 {
				printString("\n")
			}
			printString("==> " + name + " <==\n")
		}

		buf := make([]byte, n)
		bytesRead, _ := file.Read(buf)

		for j := 0; j < bytesRead; j++ {
			z01.PrintRune(rune(buf[j]))
		}

		file.Close()
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
