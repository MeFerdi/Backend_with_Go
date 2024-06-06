package main

func AtoiBase(s string, base string) int {

	baseInt := len(base)
	result := 0

	for _, char := range s {
		digit := DigitValue(char)
		result = result*baseInt + digit
	}
	return result
}

func DigitValue(char rune) int {
	if char >= '0' && char <= '9' {
		return int(char - '0')
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 'A' + 10)
	} else {
		return 0
	}
}
