package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	counts := make(map[string]int) // Create a map to store the word counts
	word := ""                     // Initialize an empty string to store the current word
	for _, char := range str {     // Iterate through each character in the input string
		if char == ' ' { // If the character is a space, it means the word is complete
			counts[word]++ // Increment the count of the word in the map
			word = ""      // Reset the word to an empty string for the next word
		} else {
			word += string(char) // Append the character to the current word
		}
	}
	counts[word]++ // Increment the count of the last word in the string
	return counts  // Return the map with the word counts
}
