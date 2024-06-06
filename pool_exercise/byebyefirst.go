package main

func ByeByeFirst(strings []string) []string {
	if len(strings) == 0 {
		return []string{}
	}
	newSlice := make([]string, len(strings)-1)

	copy(newSlice, strings[1:])

	return newSlice
}
