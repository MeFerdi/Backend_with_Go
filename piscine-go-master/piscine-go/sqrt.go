package piscine

func Sqrt(nb int) int {
	result := 0

	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			result = i
			break
		}
	}

	return result
}
