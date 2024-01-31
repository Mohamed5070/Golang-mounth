package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing\n")
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments\n")
	} else {
		fileName := os.Args[1]
		countent, _ := os.ReadFile(fileName)
		fmt.Print(string(countent))
	}
}
