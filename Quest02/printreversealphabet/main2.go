package main

import "github.com/01-edu/z01"

func main() {
	var crap rune = 'z'
	for crap >= 'a' {
		z01.PrintRune(crap)
		crap--
	}
	z01.PrintRune(10)
}
