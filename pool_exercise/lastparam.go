package main

import (
	"fmt"
	"os"
)

func main() {
	Arguments := os.Args

	fmt.Println((Arguments[len(Arguments)-1]))
}
