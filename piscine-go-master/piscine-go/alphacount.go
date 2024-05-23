package piscine

func AlphaCount(s string) int {
	Count := 0

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			Count++
		}
	}
	return Count
}
