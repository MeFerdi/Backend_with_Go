// Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

//     If the the argument is not a number, the program should print 00000000.

// Usage :

// $ go run . 1
// 00000001$
// $ go run . 192
// 11000000$
// $ go run . a
// 00000000$
// $ go run . 1 1
// $ go run .
// $
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	numb := Atoi(args)

	if numb < 0 || numb > 255 {
		numb = 0
	}

	for i := 7; i >= 0; i-- {
		bit := (numb >> i) & 1
		for _, v := range Itoa(bit) { // z01.PrintRune(rune('0' + bit)) to replace Itoa
			z01.PrintRune(v)
		}
	}
}

func Atoi(s string) int {
	var number int
	sign := 1

	for i, char := range s {
		if char == '-' && i == 0 {
			sign = -1
		} else if char == '+' && i == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return number * sign
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...)
		n /= 10
	}
	return string(digits)
}
