// Write a function that takes a byte, reverses it bit by bit (as shown in the example) and returns the result.
// Expected function

// func ReverseBits(oct byte) byte {

// }

// Example:

// 1 byte

// 0010 0110 || / 0110 0100
package main

import (
	"github.com/01-edu/z01"
)

func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = (result << 1) | (oct & 1)
		oct >>= 1
	}
	return result
}

func main() {
	z01.PrintRune(rune(ReverseBits(0o010)))
	z01.PrintRune('\n')
}
