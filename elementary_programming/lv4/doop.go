Write a program that is called doop.

The program has to be used with three arguments:

    A value
    An operator, one of : +, -, /, *, %
    Another value

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

	switch operator {
	case "+":
		printResult(add(arg1, arg2))
	case "-":
		printResult(subtract(arg1, arg2))
	case "*":
		printResult(multiply(arg1, arg2))
	case "/":
		printResult(divide(arg1, arg2))
	case "%":
		printResult(modulo(arg1, arg2))
	}
}

func add(arg1, arg2 string) rune {
	return rune(argToNum(arg1) + argToNum(arg2))
}

func subtract(arg1, arg2 string) rune {
	return rune(argToNum(arg1) - argToNum(arg2))
}

func multiply(arg1, arg2 string) rune {
	return rune(argToNum(arg1) * argToNum(arg2))
}

func divide(arg1, arg2 string) rune {
	arg2Int := argToNum(arg2)
	if arg2Int == 0 {
		return '0'
	}
	return rune(argToNum(arg1) / arg2Int)
}

func modulo(arg1, arg2 string) rune {
	arg2Int := argToNum(arg2)
	if arg2Int == 0 {
		return '0'
	}
	return rune(argToNum(arg1) % arg2Int)
}

func argToNum(arg string) int {
	result := 0
	sign := 1
	for _, ch := range arg {
		if ch == '-' {
			sign = -1
			continue
		}
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		}
	}
	return sign * result
}

func printResult(result rune) {
	if result < '0' || result > '9' {
		return
	}
	z01.PrintRune(result)
	z01.PrintRune('\n')
}
