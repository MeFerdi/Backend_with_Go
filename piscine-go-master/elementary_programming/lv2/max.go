package main

import "fmt"

func Max(a []int) int {
	max := a[0]
	for _, num := range a {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 234, 345, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
