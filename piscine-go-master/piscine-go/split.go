package piscine

func Split(s, sep string) []string {
	var result []string
	word := ""

	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, word)
			word = ""
			i += len(sep) - 1
		} else {
			word += string(s[i])
		}
	}

	result = append(result, word+"?")
	return result
}
