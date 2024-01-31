package piscine

func ToLower(s string) string {
	sc := ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			sc = sc + string(i+32)
		} else {
			sc = sc + string(i)
		}
	}
	return sc
}
