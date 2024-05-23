package piscine

func LoafOfBread(str string) string {
	MyStr1 := ""
	if str == "" {
		return "\n"
	}
	if len(str) <= 4 {
		return "Invalid Output"
	}
	n := 0
	for _, u := range str {
		if u != ' ' && n != 5 {
			MyStr1 += string(u)
			n++
		} else if n == 5 {
			MyStr1 += " "
			n = 0
		}
	}
	if len(MyStr1) > 0 && MyStr1[len(MyStr1)-1] == ' ' {
		MyStr1 = MyStr1[:len(MyStr1)-1]
	}
	return MyStr1 + "\n"
}
