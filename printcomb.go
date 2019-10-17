package piscine

import "github.com/01-edu/z01"

//var a, b, c int

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	} z01.PrintRune('\n')
}
