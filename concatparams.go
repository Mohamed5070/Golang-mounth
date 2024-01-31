package piscine

func ConcatParams(args []string) string {
	u := []string(args)
	result := ""

	for i := 0; i < len(u); i++ {
		result += u[i]
		if i < len(u)-1 {
			result += "\n"
		}
	}
	return result
}
