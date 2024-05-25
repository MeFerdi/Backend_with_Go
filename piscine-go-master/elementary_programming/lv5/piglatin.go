package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	runeMap := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	str := os.Args[1]
	var VowelInt int
	for i, c := range str {
		if runeMap[c] {
			VowelInt = i
			break
		}
	}
	if VowelInt == 0 {
		str += "ay"
	} else {
		str = str[VowelInt:] + str[:VowelInt] + "ay"
	}

	for _, res := range str {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
