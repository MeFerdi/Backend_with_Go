package main

import "github.com/01-edu/z01"

func main() {
	char := "Hello World!"

	for _, c := range char {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
