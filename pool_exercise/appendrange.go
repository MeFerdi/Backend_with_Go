package main

func AppendRange(min, max int) []int {
	var result []int

	if min >= max {
		return []int{}
	}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
