package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			f := j - 1
			for k := i; k >= '0'; k-- {
				for ; f >= '0'; f-- {
					z01.PrintRune((i))
					z01.PrintRune((j))
					z01.PrintRune(' ')
					z01.PrintRune((k))
					z01.PrintRune((f))
					if i > '0' || j > '1' || k > '0' || f > '1' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				f = '9'
			}

		}
	}
}
