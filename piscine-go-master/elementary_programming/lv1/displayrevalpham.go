package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 0 {
			z01.PrintRune(i - ('a' - 'A'))
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
