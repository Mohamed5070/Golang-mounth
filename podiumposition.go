package piscine

func PodiumPosition(podium [][]string) [][]string {
	SortStringTable(podium)
	return podium
}

func SortStringTable(table [][]string) {
	for i := 0; i < len(table); i++ {
		for j := len(table) - 1; j > i; j-- {
			if table[i][0] > table[j][0] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
