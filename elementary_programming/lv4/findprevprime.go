// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.
// Expected function

// func FindPrevPrime(nb int) int {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import "fmt"

// func main() {
// 	fmt.Println(FindPrevPrime(5))
// 	fmt.Println(FindPrevPrime(4))
// }

// And its output :

// $ go run .
// 5
// 3
// $
<<<<<<< HEAD

package main

=======
package main

import "fmt"

>>>>>>> 5176d1612acef36c1e120732bcb46e32e06b0f95
func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if IsPrime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
<<<<<<< HEAD
=======

func main() {
	fmt.Println(FindPrevPrime(7))
	fmt.Println(FindPrevPrime(6))
}
>>>>>>> 5176d1612acef36c1e120732bcb46e32e06b0f95
