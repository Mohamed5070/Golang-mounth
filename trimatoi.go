package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	index := 0
	for _, value := range s {
		if value == '-' && index == 0 {
			sign = -1
		} else if value >= '0' && value <= '9' {
			res *= 10
			res += int(value - '0')
			index++
		}
	}
	return (res * sign)
}
