package piscine

func TrimAtoi(s string) int {
	var num int
	var sign int = 1
	var found bool = false

	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			found = true
		} else if c == '-' && !found {
			sign = -1
		}
	}

	return num * sign
}
