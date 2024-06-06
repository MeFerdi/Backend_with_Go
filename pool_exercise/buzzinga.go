package main

import "fmt"

func BuzZinga(number int) {
	if number <= 0 || number%3 != 0 && number%5 != 0 {
		fmt.Println("*")
		return
	}
	if number%3 == 0 {
		fmt.Println("Buz")
	}
	if number%5 == 0 {
		fmt.Println("Zinga")
	}
	if number%3 == 0 && number%5 == 0 {
		fmt.Println("BuzZinga")
	}

}
