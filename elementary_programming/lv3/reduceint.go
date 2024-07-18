// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice, it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

// func ReduceInt(a []int, f func(int, int) int) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

// And its output :

// $ go run .
// 1000
// 502
// 250
// $
package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) < 1 {
		return
	}

	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	Itoa(result)
	z01.PrintRune('\n')
}

// println(result)
func Itoa(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		Itoa(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))

}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
