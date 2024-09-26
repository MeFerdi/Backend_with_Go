package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		fmt.Println("error: number is negative")
	}
	if n%2 == 0 && n%3 == 0 {
		fmt.Print("fish and chips")
	} else if n%2 == 0 {
		fmt.Print("fish")
	} else if n%3 == 0 {
		fmt.Print("chips")
	} else if n%2 != 0 || n%3 != 0 {
		fmt.Println("error: non divisible")
	}
	return ""
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
