package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := 0
	for i := 1; i <= len(a)-1; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
