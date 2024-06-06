package main

import "fmt"

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 8; b++ {
			t := b + 1
			for c := a; c <= 9; c++ {
				for ; t <= 9; t++ {
					fmt.Print(a)
					fmt.Print(b)
					fmt.Print(" ")
					fmt.Print(c)
					fmt.Print(t)
					if a < 9 || b < 8 || c < 9 || t < 9 {
						fmt.Print(", ")
					}

				}
				t = 0
			}
		}

	}
	fmt.Println()
}
func main() {
	PrintComb2()
}
