package main

import "fmt"

func PrintNbr(n int) {

	if n < -123456789 || n > 123456789 {
		fmt.Print("error")
		return
	}
	fmt.Print(n)
}
