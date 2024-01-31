package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	i := l
	j := 0
	for i != nil {
		if pos == j {
			return i
		}
		j++
		i = i.Next
	}
	return nil
}
