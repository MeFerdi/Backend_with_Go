package piscine

func SortIntegerTable(table []int) {
	for k := 0; k < len(table); k++ {
		for l := k + 1; l < len(table); l++ {
			if table[k] > table[l] {
				table[k], table[l] = table[l], table[k]
			}
		}
	}
}
