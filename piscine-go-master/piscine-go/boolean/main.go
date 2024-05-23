package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	EvenMsg = "I have an even number of arguments\n"
	OddMsg  = "I have an odd number of arguments\n"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func lengthOfArg(args []string) int {
	return len(args)
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(lengthOfArg(os.Args[1:])) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
