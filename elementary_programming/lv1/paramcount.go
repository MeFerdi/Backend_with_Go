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
	Args := os.Args[1:]
	count := 0

	for i := 0; i < len(Args); i++ {
		count++
	}
	z01.PrintRune(rune(count) + '0')
	z01.PrintRune('\n')
}
