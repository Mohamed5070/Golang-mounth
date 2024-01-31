package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	str := []rune(Itoa(n))
	Sort(str)
	PrintS(string(str))
}

func Itoa(n int) string {
	var str string
	if n == 0 {
		str = "0"
	}
	for n > 0 {
		str = str + string(rune(n%10)+'0')
		n = n / 10
	}
	return str
}

func Sort(n []rune) string {
	for i := 0; i < len(n); i++ {
		for j := len(n) - 1; j > i; j-- {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return string(n)
}

func PrintS(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
