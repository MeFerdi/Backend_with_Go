package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	Args := os.Args[1:]
	for i := 0; i < len(Args); i++ {
		count++
	}
	if count <= len(Args) {
		var str string
		for count > 0 {
			str = string('0'+count%10) + str
			count /= 10
		}
		for _, res := range str {
			z01.PrintRune(rune(res))
		}

	}
	z01.PrintRune('\n')
}
