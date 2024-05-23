package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(str string) {
	arrayStr := []rune(str)
	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func PrintValue(a int) {
	r := '0'
	if a/10 > 0 {
		PrintValue(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintValue(points.x)
	PrintStr(", y = ")
	PrintValue(points.y)
	PrintStr("\n")
}
