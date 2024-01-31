package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for n := 1; n < len(args); n++ {
		d := BAtoi(args[n])
		if d >= 1 && d <= 26 {
			z01.PrintRune(rune(d+'a') - 1)
		}
	}
	z01.PrintRune('\n')
}

func BAtoi(s string) int {
	res := 0
	for _, n := range s {
		res = res*10 + int(n-'0')
	}
	return res
}
