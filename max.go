package main

import "fmt"

func Max(a []int) int {
	max := 0
	if len(a) == 0 {
		return 0
	}
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 254, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
