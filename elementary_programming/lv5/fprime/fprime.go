package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 2 {
		return
	}

	primeFactors(num)
}

func primeFactors(n int) {
	first := true
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			n /= i
			first = false
		}
	}
	fmt.Println()
}
