// Write a program that displays a number's multiplication table.

//     The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.

// Usage

// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		return
	}

	for i := 1; i <= 9; i++ {
		result := num * i
		printMultiplication(i, num, result)
		z01.PrintRune('\n')
	}
}

func printMultiplication(a, b, result int) {
	printNumber(a)
	z01.PrintRune(' ')
	z01.PrintRune('x')
	z01.PrintRune(' ')
	printNumber(b)
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(result)
}

func printNumber(n int) {
	if n >= 10 {
		printNumber(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
