package piscine

func ToUpper(s string) string {
	sc := ""
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			sc = sc + string(i-32)
		} else {
			sc = sc + string(i)
		}
	}
	return sc
}
