package piscine

func Compare(a, b string) int {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	if len(a) == 0 {
		return -1
	}
	if len(b) == 0 {
		return 1
	}
	if a[0] < b[0] {
		return -1
	}
	if a[0] > b[0] {
		return 1
	}
	return Compare(a[1:], b[1:])
}
