package piscine

func JumpOver(str string) string {
	s := ""

	for i := 1; i <= len(str); i++ {
		if i%3 == 0 {
			s += string(str[i-1])
		}
	}
	s += "\n"
	return s
}
