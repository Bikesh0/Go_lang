package main

import (
	"os"
)

func printString(s string) {
	os.Stdout.Write([]byte(s))
}

func atoi(s string) (int, bool) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		return
	}

	n, ok := atoi(args[1])
	if !ok || n < 0 {
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

		if len(files) > 1 {
			if i > 0 {
				printString("\n")
			}
			printString("==> " + name + " <==\n")
		}

		buf := make([]byte, n)
		bytesRead, _ := file.Read(buf)

		os.Stdout.Write(buf[:bytesRead])

		file.Close()
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
