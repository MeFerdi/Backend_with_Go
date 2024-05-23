package piscine

func ToLower(s string) string {
	set := ""

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		set += string(char)
	}
	return set
}
