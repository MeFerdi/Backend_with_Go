package piscine

func StrRev(s string) string {
	char := []rune(s)
	var result []rune

	for i := len(char) - 1; i >= 0; i-- {
		result = append(result, char[i])
	}
	return (string(result))
}
