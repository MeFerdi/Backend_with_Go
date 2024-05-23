package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	MyStr := []rune(s)
	for _, val := range MyStr {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
