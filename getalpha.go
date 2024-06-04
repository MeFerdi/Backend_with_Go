package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println()
		return
	}

	if num < 0 || num > 127 {
		fmt.Println()
		return
	}

	character := string(num)
	fmt.Println(character)
}
