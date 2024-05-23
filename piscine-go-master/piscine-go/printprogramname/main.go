package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := getProgramName(os.Args[0])
	for _, r := range programName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func getProgramName(path string) string {
	var programName string
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			programName = path[i+1:]
			break
		}
	}
	return programName
}
