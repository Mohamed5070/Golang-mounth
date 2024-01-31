package piscine

func Compact(ptr *[]string) int {
	str := *ptr
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != "" {
			str[count] = str[i]
			count++
		}
	}
	*ptr = str[:count]
	return count
}
