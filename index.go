package piscine

func Index(s string, toFind string) int {
	strToFind := []rune(toFind)
	str := []rune(s)
	for i := 0; i < len(str)-len(strToFind); i++ {
		if Compare(string(str[i:i+len(strToFind)]), string(strToFind)) == 0 {
			return i
		}
	}
	return -1
}
