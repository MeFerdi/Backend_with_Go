package piscine

func UltimateDivMod(a *int, b *int) {
	s := *a
	*a = *a / *b
	*b = s % *b
}
