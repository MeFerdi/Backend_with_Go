// Write a program that transforms a string passed as argument in its Pig Latin version.

// The rules used by Pig Latin are as follows:

//     If a word begins with a vowel, just add "ay" to the end.
//     If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
//     Only the latin vowels will be considered as vowel(aeiou).

// If the word has no vowels, the program should print "No vowels".

// If the number of arguments is different from one, the program prints nothing.
// Usage

// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $

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
