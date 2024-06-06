package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var snakeCase string
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				snakeCase += "_"
			}
			snakeCase += string(r + 32) // convert to lowercase
		} else {
			snakeCase += string(r)
		}
	}
	return snakeCase
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Output: Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // Output: hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // Output: camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // Output: CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // Output: camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // Output: hey2
}
