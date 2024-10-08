// Write a program called rn. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

// The program should have a limit of 4000. In case of an invalid number, for example "hello" or 0 the program should print ERROR: cannot convert to roman digit.

// Roman Numerals reminder:
// I 	1
// V 	5
// X 	10
// L 	50
// C 	100
// D 	500
// M 	1000

// For example, the number 1732 would be denoted MDCCXXXII in Roman numerals. However, Roman numerals are not a purely additive number system. In particular, instead of using four symbols to represent a 4, 40, 9, 90, etc. (i.e., IIII, XXXX, VIIII, LXXXX, etc.), such numbers are instead denoted by preceding the symbol for 5, 50, 10, 100, etc., with a symbol indicating subtraction. For example, 4 is denoted IV, 9 as IX, 40 as XL, etc.

// The following table gives the Roman numerals for the first few positive integers.
// 1 	I 	11 	XI 	21 	XXI
// 2 	II 	12 	XII 	22 	XXII
// 3 	III 	13 	XIII 	23 	XXIII
// 4 	IV 	14 	XIV 	24 	XXIV
// 5 	V 	15 	XV 	25 	XXV
// 6 	VI 	16 	XVI 	26 	XXVI
// 7 	VII 	17 	XVII 	27 	XXVII
// 8 	VIII 	18 	XVIII 	28 	XXVIII
// 9 	IX 	19 	XIX 	29 	XXIX
// 10 	X 	20 	XX 	30 	XXX
// Usage

// $ go run . hello
// ERROR: cannot convert to roman digit
// $ go run . 123
// C+X+X+I+I+I
// CXXIII
// $ go run . 999
// (M-C)+(C-X)+(X-I)
// CMXCIX
// $ go run . 3999
// M+M+M+(M-C)+(C-X)+(X-I)
// MMMCMXCIX
// $ go run . 4000
// ERROR: cannot convert to roman digit
// $
package main

<<<<<<< HEAD
=======
package main

>>>>>>> refs/remotes/origin/main
import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
<<<<<<< HEAD
	const Error = "ERROR: cannot convert to roman digit"

	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	var result []string
	var result1 []string

	number := Atoi(args)
	if number == 0 || number >= 4000 {
		for _, c := range Error {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}

	for number > 0 {
		if number >= 1000 {
			result1, result = append(result1, "M"), append(result, "M")
			number -= 1000
		} else if number >= 900 {
			result1, result = append(result1, "CM"), append(result, "(M-C)")
			number -= 900
		} else if number >= 500 {
			result1, result = append(result1, "D"), append(result, "D")
			number -= 500
		} else if number >= 400 {
			result1, result = append(result1, "CD"), append(result, "(D-C)")
			number -= 400
		} else if number >= 100 {
			result1, result = append(result1, "C"), append(result, "C")
			number -= 100
		} else if number >= 90 {
			result1, result = append(result1, "XC"), append(result, "(C-X)")
			number -= 90
		} else if number >= 50 {
			result1, result = append(result1, "L"), append(result, "L")
			number -= 50
		} else if number >= 40 {
			result1, result = append(result1, "XL"), append(result, "(L-X)")
			number -= 40
		} else if number >= 10 {
			result1, result = append(result1, "X"), append(result, "X")
			number -= 10
		} else if number >= 9 {
			result1, result = append(result1, "IX"), append(result, "(X-I)")
			number -= 9
		} else if number >= 5 {
			result1, result = append(result1, "V"), append(result, "V")
			number -= 5
		} else if number >= 4 {
			result1, result = append(result1, "IV"), append(result, "(v-I)")
			number -= 4
		} else {
			result1, result = append(result1, "I"), append(result, "I")
			number -= 1
		}
	}
	for i, c := range result {
		for _, n := range c {
			z01.PrintRune(n)
		}
		if i < len(result)-1 {
			z01.PrintRune('+')
		}
	}
	z01.PrintRune('\n')
	for _, n := range result1 {
		for _, c := range n {
			z01.PrintRune(c)
		}
=======
	if len(os.Args) != 2 {
		PrintError()
		return
	}

	arg := os.Args[1]
	num := Atoi(arg)
	if num <= 0 || num >= 4000 {
		PrintError()
		return
	}

	roman, calculation := ToRoman(num)
	PrintStr(calculation)
	PrintStr(roman)
}

func ToRoman(num int) (string, string) {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman, calculation := "", ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += sym[i]
			if calculation != "" {
				calculation += "+"
			}
			calculation += sym[i]
		}
	}
	return roman, calculation
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
>>>>>>> refs/remotes/origin/main
	}
	z01.PrintRune('\n')
}

<<<<<<< HEAD
// strconv.Atoi is allowed
func Atoi(s string) int {
	var number int
	sign := 1

	for i, char := range s {
		if char == '-' && i == 0 {
			sign = -1
		} else if char == '+' && i == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		}
	}
	return number * sign
=======
func PrintError() {
	msg := "ERROR: cannot convert to roman digit"
	PrintStr(msg)
}

func Atoi(s string) int {
	q := 0
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			q = q*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return q * sign
>>>>>>> refs/remotes/origin/main
}
