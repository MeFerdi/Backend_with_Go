package piscine

func Swap(a *int, b *int) {
	tempA := *b
	tempB := *a
	*b = tempB
	*a = tempA
}
