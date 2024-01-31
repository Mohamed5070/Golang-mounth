package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, ascii := range str {
		result = append(result, int(ascii))
	}
	return result
}
