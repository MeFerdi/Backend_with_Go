// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

//     If the number of arguments is different from 3, the program displays nothing.

//     If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage

// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
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
	Args := os.Args[1:]
	var words string
	for _, arg := range Args[0] {
		if string(arg) == Args[1] {
			words += Args[2]
		} else {
			words += string(arg)
		}
	}
	for _, res := range words {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
