package piscine

func LastRune(s string) rune {
	sc := []rune(s)
	return sc[len(sc)-1]
}
