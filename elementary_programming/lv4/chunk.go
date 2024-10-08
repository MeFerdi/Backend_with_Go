// Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

//     If the size is 0 it should print a newline ('\n').

// Expected function

// func Chunk(slice []int, size int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }

// And its output :

// $ go run .
// []

// [[0 1 2] [3 4 5] [6 7]]
// [[0 1 2 3 4] [5 6 7]]
// [[0 1 2 3] [4 5 6 7]]
// $
<<<<<<< HEAD

package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
=======
package main

import "fmt"

func Chunk(slice []int, size int) {
	count := 0
	var newSlice []int
	var outputSlice [][]int

>>>>>>> 5176d1612acef36c1e120732bcb46e32e06b0f95
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
<<<<<<< HEAD

	outputSlice := [][]int{}
	for size < len(slice) {
		outputSlice = append(outputSlice, slice[:size])
		slice = slice[size:]
	}
	outputSlice = append(outputSlice, slice)

	fmt.Println(outputSlice)
}
=======
	for _, char := range slice {
		newSlice = append(newSlice, char)
		count++
		if count == size {
			outputSlice = append(outputSlice, newSlice)
			newSlice = []int{}
			count = 0
		}
	}
	if len(newSlice) > 0 {
		outputSlice = append(outputSlice, newSlice)
	}
	fmt.Println(outputSlice)
}
func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
>>>>>>> 5176d1612acef36c1e120732bcb46e32e06b0f95
