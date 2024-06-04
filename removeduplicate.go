package main

func RemoveDuplicate(slice []int) []int {

	unique := make(map[int]bool)
	result := []int{}

	for _, num := range slice {

		if _, ok := unique[num]; !ok {

			unique[num] = true
			result = append(result, num)
		}
	}

	return result
}
