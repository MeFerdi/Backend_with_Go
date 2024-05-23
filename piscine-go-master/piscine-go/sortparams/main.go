package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sortASCII(args)
	for _, arg := range args {
		printString(arg)
		printChar('\n')
	}
}

func sortASCII(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compareASCII(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compareASCII(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return len(a) - len(b)
}

func printString(s string) {
	for _, r := range s {
		printChar(r)
	}
}

func printChar(r rune) {
	z01.PrintRune(r)
}
