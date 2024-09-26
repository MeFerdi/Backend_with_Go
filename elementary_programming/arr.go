package main

import (
	"fmt"
)

func main() {
	arr := []int{}
	for i := 0; i <= 15; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
}
