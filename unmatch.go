package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 0
		k := a[i]
		for j := 0; j < len(a); j++ {
			if k == a[j] {
				count++
			}
		}
		if count%2 == 1 {
			return k
		}
	}
	return -1
}
