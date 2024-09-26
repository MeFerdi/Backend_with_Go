package main

func CamelToSnakeCase(s string) string {
	word := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' && i > 0 && i < len(s)-1 {
			word += "_"
			word += string(char)
		} else {
			word += string(char)
		}
	}
	return word
}

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }
