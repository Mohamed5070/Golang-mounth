package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	slice := []string{}
	for _, i := range s {
		if i != ' ' {
			str += string(i)
		} else if str != "" {
			slice = append(slice, str)
			str = ""
		}
	}
	if str != "" {
		slice = append(slice, str)
	}

	return slice
}
