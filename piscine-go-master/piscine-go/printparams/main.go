package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		printString(arg)
		z01.PrintRune('\n')
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
