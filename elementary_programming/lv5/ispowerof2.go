// Write a program that determines if a given number is a power of 2. A number n is a power of 2 if it meets the following condition: n = 2 m where m is another integer number.

// For example, 4 = 2 2 or 16 = 2 4 are power of 2 numbers while 24 is not.

// This program must print true if the given number is a power of 2, otherwise it has to print false.

//     If there is more than one or no argument, the program should print nothing.

//     When there is only one argument, it will always be a positive valid int.

// Usage :

// $ go run . 2 | cat -e
// true$
// $ go run . 64
// true
// $ go run . 513
// false
// $ go run .
// $ go run . 64 1024
// $
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
<<<<<<< HEAD

	arg := os.Args[1]
	newArg := 0
	for _, v := range arg {
		n := int(v - 48)

		newArg = (newArg * 10) + n

	}

=======
	arg := os.Args[1]
	newArg := 0
	for _, char := range arg {
		n := int(char - 48)
		newArg = (newArg * 10) + n
	}
>>>>>>> refs/remotes/origin/main
	fmt.Println(isPowerof2(newArg))
}

func isPowerof2(newArg int) bool {
	return newArg%2 == 0
}
