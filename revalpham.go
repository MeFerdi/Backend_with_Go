package main

import "fmt"

func main() {
	for char := 'z'; char >= 'a'; char-- {
		if char%2 == 0 {
			fmt.Print(string(char - ('a' - 'A')))
		} else {
			fmt.Print(string(char))
		}
	}
	fmt.Println()
}
