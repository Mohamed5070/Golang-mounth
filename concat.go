package piscine

func Concat(str1 string, str2 string) string {
	s3 := ""
	for i := 0; i < len(str1); i++ {
		s3 = s3 + string(str1[i])
	}
	for i := 0; i < len(str2); i++ {
		s3 = s3 + string(str2[i])
	}
	return s3
}
