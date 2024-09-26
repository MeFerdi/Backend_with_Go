package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	var result int
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
		if count > 0 {
			result += count
			count = 0
		}
	}
	return result
}
func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
