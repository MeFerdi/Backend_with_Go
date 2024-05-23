package piscine

type predicate func(string) bool

func CountIf(f predicate, tab []string) int {
	count := 0
	for _, s := range tab {
		if f(s) {
			count++
		}
	}
	return count
}

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
