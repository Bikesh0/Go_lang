package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		return
	}

	n := 0
	for _, c := range args[1] {
		if c < '0' || c > '9' {
			return
		}
		n = n*10 + int(c-'0')
	}

	files := args[2:]
	hasError := false

	for i, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Println(err)
			hasError = true
			continue
		}

		info, _ := file.Stat()
		size := info.Size()

		start := int64(0)
		if int64(n) < size {
			start = size - int64(n)
		}

		file.Seek(start, 0)

		// Header for multiple files
		if len(files) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", name)
		}

		buf := make([]byte, n)
		bytesRead, _ := file.Read(buf)

		fmt.Print(string(buf[:bytesRead]))

		file.Close()
	}

	if hasError {
		os.Exit(1)
	}
}
