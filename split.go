package piscine

func Split(s, sep string) []string {
	var result []string
	Separator := false
	c := 0

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if !Separator {
				result = append(result, s[c:i])
			}
			Separator = true
			c = i + len(sep)
			i += len(sep) - 1
		} else {
			Separator = false
		}
	}

	if !Separator {
		result = append(result, s[c:])
	}

	return result
}
