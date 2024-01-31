package main

import "github.com/01-edu/z01"

func main() {
	str := "x = ab, y = bd\n"

	for _, i := range str {
		if i == 'a' {
			z01.PrintRune(52)
		} else if i == 'b' {
			z01.PrintRune(50)
		} else if i == 'd' {
			z01.PrintRune(49)
		} else {
			z01.PrintRune(i)
		}
	}
}
