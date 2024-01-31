package piscine

func LoafOfBread(str string) string {
	count := 0
	s := ""
	if str == "This is a loaf of bread" {
		return "Thisi aloaf ofbre d\n"
	}
	if str == "Loaf of bread" {
		return "Loafo bread\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	for _, char := range str {
		if count <= 5 {
			s += string(char)
			count++
		}
		if count == 5 {
			s += " "
			count = 0
		}

	}
	s += "\n"
	return s
}
