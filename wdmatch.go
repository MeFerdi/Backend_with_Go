package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is not equal to 2
	if len(os.Args) != 3 {
		return
	}

	// Get the two input strings
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Check if it is possible to write the first string with characters from the second string
	if isPossible(str1, str2) {
		// Display the first string followed by a newline
		fmt.Println(str1)
	}
}

func isPossible(str1, str2 string) bool {
	// Initialize index variables for both strings
	i := 0
	j := 0

	// Iterate through both strings
	for i < len(str1) && j < len(str2) {
		// If the characters match, move to the next character in both strings
		if str1[i] == str2[j] {
			i++
		}
		// Move to the next character in the second string
		j++
	}

	// If all characters in the first string are found in the second string in order, return true
	return i == len(str1)
}
