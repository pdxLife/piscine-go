package main

import "github.com/01-edu/z01"

func main() {
	i := 48
	for i <= 57 {
		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune(10)
}
