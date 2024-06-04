package main

import (
	"fmt"
	"os"
	// "github.com/01-edu/z01"
)

func main() {
	count := 0
	Args := os.Args

	for i := 1; i < len(Args); i++ {
		count++
	}
	if count <= len(Args) {
		fmt.Println(count)
	}
}
