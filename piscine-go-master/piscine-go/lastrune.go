package piscine

func LastRune(s string) rune {
	MyVar := []rune(s)

	return MyVar[len(MyVar)-1]
}
