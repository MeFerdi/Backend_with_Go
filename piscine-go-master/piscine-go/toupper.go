package piscine

func ToUpper(s string) string {
	set := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 32
		}
		set += string(char)
	}
	return set
}
