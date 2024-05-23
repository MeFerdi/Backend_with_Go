// wdmatch
// Instructions
// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

// If the number of arguments is different from 2, the program displays nothing.

// Usage
// $ go run . 123 123
// 123
// $ go run . faya fgvvfdxcacpolhyghbreda
// faya
// $ go run . faya fgvvfdxcacpolhyghbred
// $ go run . error rrerrrfiiljdfxjyuifrrvcoojh
// $ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
// quarante deux
// $ go run .
// $
package main

import (
	"fmt"
	"os"
)

func WordMatch(str1, str2 string) bool {
	if len(str1) == 0 || len(str2) == 0 {
		return false
	}

	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	return i == len(str1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1, str2 := os.Args[1], os.Args[2]

	if WordMatch(str1, str2) {
		fmt.Println(str1)
	}
}
