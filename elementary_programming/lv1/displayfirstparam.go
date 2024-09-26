package main

import (
	"os"

	"github.com/01-edu/z01"
)

func FirstParam() {
	if len(os.Args) < 2 {
		return
	}
	Args := os.Args[1]
	for _, char := range Args {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
