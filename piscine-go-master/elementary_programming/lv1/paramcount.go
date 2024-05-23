package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	count := 0
	for i := 0; i < len(Args); i++ {
		count++
	}
	z01.PrintRune(rune(count) + '0')
	z01.PrintRune('\n')
}
