package piscine

func Rot14(s string) string {
	str := ""
	for _, c := range s {
		if (c >= 'a' && c <= 'l') || (c >= 'A' && c <= 'L') {
			str += string(c + 14)
		} else if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			str += string(c - 12)
		} else if c == 'm' {
			str += string('a')
		} else if c == 'M' {
			str += string('A')
		} else {
			str += string(c)
		}
	}
	return str
}
