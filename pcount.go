package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	z01.PrintRune(rune(len(Args)))
}
