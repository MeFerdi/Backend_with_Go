package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[len(os.Args)-1]
	for _, char := range Args {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
