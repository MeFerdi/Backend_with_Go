package piscine

func NRune(s string, n int) rune {
	Count := 0

	for _, v := range s {

		Count++

		if Count == n {
			return v
		}
	}
	return 0
}
