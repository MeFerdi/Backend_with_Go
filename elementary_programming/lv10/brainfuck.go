package main

import (
	"fmt"
)

func interpretBrainfuck(sourceCode string) {
	memory := make([]byte, 2048)
	pointer := 0

	stack := make([]int, 0)
	jumps := make(map[int]int)

	// Find matching brackets and store their positions
	for i, char := range sourceCode {
		if char == '[' {
			stack = append(stack, i)
		} else if char == ']' {
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jumps[start] = i
			jumps[i] = start
		}
	}

	for i := 0; i < len(sourceCode); i++ {
		command := sourceCode[i]

		switch command {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			fmt.Printf("%c", memory[pointer])
		case '[':
			if memory[pointer] == 0 {
				i = jumps[i]
			}
		case ']':
			if memory[pointer] != 0 {
				i = jumps[i]
			}
		}
	}

	fmt.Println()
}
