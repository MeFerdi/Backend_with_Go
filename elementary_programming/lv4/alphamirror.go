// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

// The case of the letter remains unchanged, for example :

// 'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

// The final result will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program prints a new line.
// Usage

// $ go run . "abc"
// zyx$
// $
// $ go run . "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// $
// $ go run .
// $

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	output := ""
	for _, char := range input {
		if isAlphabetical(char) {
			if isLower(char) {
				output += string('z' - (char - 'a'))
			} else {
				output += string('Z' - (char - 'A'))
			}
		} else {
			output += string(char)
		}
	}

	fmt.Println(output)
}

func isAlphabetical(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isLower(char rune) bool {
	return char >= 'a' && char <= 'z'
}
