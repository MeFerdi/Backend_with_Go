package main

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if isAlphab(char) {
			count++
		}
	}
	return count
}

func isAlphab(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

// func main() {
// 	str := "Hello 123"
// 	res := CountAlpha(str)
// 	fmt.Println(res)
// }
