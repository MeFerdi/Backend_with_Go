package piscine

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func medianOfFive(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
	return arr[2]
}

func Abort(a, b, c, d, e int) int {
	return medianOfFive(a, b, c, d, e)
}
