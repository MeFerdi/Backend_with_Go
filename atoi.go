package main

func Atoi(s string) int {
	result := 0
	sign := 1

	for i, char := range s {

		if char == '+' && i == 0 { //assigning the appropriate sign to the values.
			continue
		} else if char == '-' && i == 0 {
			sign = -1
			continue
		}
		if char < '0' || char > '9' {
			return 0
		}

		result = result*10 + int(char-'0') //converting the char value to int, by getting the difference in the ASCII values.
	}

	return result * sign
}
