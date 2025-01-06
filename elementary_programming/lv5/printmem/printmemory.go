package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i += 4 {

		for j := i; j < i+4 && j < len(arr); j++ {
			hexPrint(arr[j])
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')

		for j := i; j < i+4 && j < len(arr); j++ {
			if arr[j] >= 32 && arr[j] <= 126 {
				z01.PrintRune(rune(arr[j]))
			} else {
				z01.PrintRune('.')
			}
		}
		z01.PrintRune('\n')
	}
}

func hexPrint(b byte) {
	hexChars := "0123456789abcdef"
	z01.PrintRune(rune(hexChars[b>>4]))
	z01.PrintRune(rune(hexChars[b&0x0F]))
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
