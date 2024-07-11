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

	a := Atoi(os.Args[1])
	// if err != nil {
	// 	return
	// }

	op := os.Args[2]
	if op != "+" && op != "-" && op != "*" && op != "/" && op != "%" {
		return
	}

	b := Atoi(os.Args[3])
	// if err != nil {
	// 	return
	// }

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			char := "No division by 0"
			for _, res := range char {
				z01.PrintRune(res)
			}
			z01.PrintRune('\n')
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			char := "NO modulo by 0"
			for _, res := range char {
				z01.PrintRune(res)
			}
			z01.PrintRune('\n')
			return
		}
		result = a % b
	}

	Itoa(result)
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	var n int
	var sign int = 1
	for i, c := range s {
		if i == 0 && (c == '-' || c == '+') {
			if c == '-' {
				sign = -1
			}
		} else if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return n * sign
}

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
