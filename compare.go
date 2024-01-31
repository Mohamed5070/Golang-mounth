package piscine

func Compare(a, b string) int {
	s1 := []rune(a)
	s2 := []rune(b)
	SL := 0
	if a < b {
		SL = len(s1)
	} else {
		SL = len(s2)
	}
	for i := 0; i < SL; i++ {
		if s1[i] < s2[i] {
			return -1
		} else if s1[i] > s2[i] {
			return 1
		}
	}
	if len(s1) < len(s2) {
		return -1
	} else if len(s1) > len(s2) {
		return 1
	} else {
		return 0
	}
}
