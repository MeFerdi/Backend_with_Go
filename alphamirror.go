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
