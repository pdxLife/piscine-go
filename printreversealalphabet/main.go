package main

import "github.com/01-edu/z01"

func main() {
	i := 122
	for i > 96 {
		z01.PrintRune(rune(i))
		i--
	}
	z01.PrintRune('\n')
}
