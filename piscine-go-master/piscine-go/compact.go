package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	compact := original[:0]
	count := 0
	for _, s := range original {
		if s != "" {
			compact = append(compact, s)
			count++
		}
	}
	*ptr = compact
	return count
}
