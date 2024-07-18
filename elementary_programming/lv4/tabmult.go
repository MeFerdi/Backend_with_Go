package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	q := 0
	for _, r := range s {
		q = q*10 + int(r-'0')
	}
	return q
}

<<<<<<< HEAD
// $
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)
=======
func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := ""
	if n < 0 {
		n = -n
		sign = "-"
	}
	q := ""
	for n > 0 {
		digits := n % 10
		q = string(rune('0'+digits)) + q
		n /= 10
	}
	return sign + q
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
>>>>>>> refs/remotes/origin/main

func main() {
	if len(os.Args) != 2 {
		return
	}
<<<<<<< HEAD

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
=======
	args := os.Args[1]
	num := BasicAtoi(args)

	for i := 1; i <= 9; i++ {
		table := num * i
		PrintStr(Itoa(i) + " x " + Itoa(num) + " = " + Itoa(table))
	}
}
>>>>>>> refs/remotes/origin/main
