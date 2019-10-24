package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			for a := '0'; a <= '9'; a++ {
				for b := '0'; b <= '9'; b++ {
					if i <= j {
						if j <= a {
							if a <= b {

							}
						}
					}
				}
			}
			if i < j {
				if j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}

		}
	}
	// z01.PrintRune('7')
	// z01.PrintRune('8')
	// z01.PrintRune('9')
	// z01.PrintRune('\n')

}
