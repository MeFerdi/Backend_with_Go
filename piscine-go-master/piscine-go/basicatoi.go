package piscine

func BasicAtoi(s string) int {
	var result int

	for _, char := range s {
		digit := char - '0'
		result = result*10 + int(digit)
	}
	return result
}
