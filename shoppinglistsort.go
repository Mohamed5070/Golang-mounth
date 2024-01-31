package piscine

func ShoppingListSort(slice []string) []string {
	SortI(slice)
	return slice
}

func SortI(table []string) {
	for i := 0; i < len(table); i++ {
		for j := len(table) - 1; j > i; j-- {
			if len(table[i]) > len(table[j]) {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
