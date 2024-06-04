// package main
// import(
// 	"github.com/01-edu/z01"
// 	"os"
// )
// func ReduceInt(a []int, f func(int, int) int) {
// 	if len(os.Args) != 1{
// 		return
// 	}
// 	result := a[0]
// 	for i := 1; i < len(a); i++ {
//         result = f(result, a[i])
//     }
// 	func Add(a int, b int)int{
// 		return a + b
// 	}
// 	func Mul(a int, b int)int{
// 		return a * b
// 	}
// 	func Modul(a int, b int)int{
// 		return a / b
// 	}

// }
// func Itoa(n int) string {
// 	if n < 0{
// 		return
// 	}
// 	for _, num
// }
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
package main

import (
	"github.com/01-edu/z01"
)

func printNumber(num int) {
	if num == 0 {
		z01.PrintRune('0')
	} else {
		digits := []int{}
		for num != 0 {
			digit := num % 10
			digits = append(digits, digit)
			num /= 10
		}

		for i := len(digits) - 1; i >= 0; i-- {
			if i == len(digits)-1 && digits[i] == 0 {
				continue
			}
			z01.PrintRune(rune('0' + digits[i]))
		}

	}
}

func main() {
	s := 056
	printNumber(s)
	z01.PrintRune('\n')
}
