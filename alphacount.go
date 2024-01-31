package piscine

func AlphaCount(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {
			j++
		}
	}
	return j
}
