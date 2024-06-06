package main

func RecursivePower(nb int, power int) int {
	var result int

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	result = nb * RecursivePower(nb, power-1)
	return result
}
