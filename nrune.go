package piscine

func NRune(s string, n int) rune {
	sc := []rune(s)
	if n <= 0 || n > len(sc) {
		return 0
	}
	return sc[n-1]
}
