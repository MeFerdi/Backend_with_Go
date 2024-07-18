// Write a program that is called doop.

// The program has to be used with three arguments:

//     A value
//     An operator, one of : +, -, /, *, %
//     Another value

// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.
// Usage

// $ go run .
// $ go run . 1 + 1 | cat -e
// 2
// $
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0
// $
// $ go run . 1 % 0 | cat -e
// No modulo by 0
// $
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[3]
	operator := os.Args[2]

	a1 := Atoi(arg1)

	a2 := Atoi(arg2)

	switch operator {
	case "+":
		result := add(a1, a2)
		Itoa(result)
	case "-":
		result := subtract(a1, a2)
		Itoa(result)
	case "*":
		result := multiply(a1, a2)
		Itoa(result)
	case "/":
		if a2 == 0 {
			char := "No division by 0"
			for _, res := range char {
				z01.PrintRune(res)
			}
			return
		}
		result := divide(a1, a2)
		Itoa(result)
	case "%":
		if a2 == 0 {
			char := "No modulo by 0"
			for _, res := range char {
				z01.PrintRune(res)
			}
			return
		}
		result := modulo(a1, a2)
		Itoa(result)
	}
}

func Atoi(s string) int {
	result := 0
	sign := 1
	for _, ch := range s {
		if ch == '-' {
			sign = -1
			continue
		}
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return sign * result
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func modulo(a, b int) int {
	return a % b
}

func Itoa(result int) {
	if result < 0 {
		z01.PrintRune('-')
		result = -result
	}
	if result >= 10 {
		Itoa(result / 10)
	}
	z01.PrintRune(rune(result%10 + '0'))
	z01.PrintRune('\n')
}
