package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var slice1 []int = make([]int, 10)

	// Generate Fibonacci numbers for multiples of 5
	for i := 0; i < len(slice1); i++ {
		slice1[i] = fibonacci(i)
	}

	// Print the slice with Fibonacci numbers
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Fibonacci number at index %d is %d\n", i, slice1[i])
	}

	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
