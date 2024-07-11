// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func ParamCount() {
// 	Args := os.Args[1:]
// 	count := 0
// 	for i := 0; i < len(Args); i++ {
// 		count++
// 	}
// 	z01.PrintRune(rune(count) + '0')
// 	z01.PrintRune('\n')
// }

// Alternative of solving the question in a way to handle large numbers
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := 0
	Args := os.Args[1:]
	for i := 0; i < len(Args); i++ {
		count++
	}
	if count <= len(Args) {
		var str string
		for count > 0 {
			str = string(count%10+'0') + str
			count /= 10
		}
		for _, res := range str {
			z01.PrintRune(res)
		}

	}
	z01.PrintRune('\n')
}
