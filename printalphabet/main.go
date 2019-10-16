package main

import "github.com/01-edu/z01"

func main() {
	i := 97
	for i < 123 {
		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune('\n')
}
