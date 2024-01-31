package piscine

func StrRev(s string) string {
	sc := []byte(s)
	i := 0
	j := len(s) - 1

	for i < j {
		sc[i], sc[j] = sc[j], sc[i]
		i++
		j--
	}
	return string(sc)
}
