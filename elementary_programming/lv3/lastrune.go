package main

func LastRune(s string) rune {
	return []rune(s)[len([]rune(s))-1]
}
