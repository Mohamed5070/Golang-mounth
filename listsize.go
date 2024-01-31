package piscine

func ListSize(l *List) int {
	count := 0
	i := l.Head
	for i != nil {
		count++
		i = i.Next
	}
	return count
}
