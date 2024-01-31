package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for _, char := range args[2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
