package piscine

func Join(strs []string, sep string) string {
	var result string

	count := 0

	for range strs {
		count++
	}
	for i := 0; i < count-1; i++ {
		result += strs[i] + sep
	}
	if count > 0 {
		result += strs[count-1]
	}
	return result
}
