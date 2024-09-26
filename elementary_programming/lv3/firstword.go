package main

func FirstWord(s string) string {
	result := ""
	for i := range s {
		if i != 0 && s[i] == ' ' {
			result += string(s[:i])
			break
		}
	}
	return result + "\n"
}

// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("hello   .........  bye"))
// }
