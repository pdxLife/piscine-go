package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for a := '0'; a <= '9'; a++ {
				for b := '0'; b <= '9'; b++ {
					if i == a && j < b || i < a {
						// if i <= j {
						// 	if j <= a {
						// 		if a <= b {
						// if i == 0 && j == 0 && a == 0 && b == 0 || i == 9 && j == 9 && a == 9 && b == 9 {
						// z01.PrintRune('')
						// } else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(a)
						z01.PrintRune(b)
						if i != '9' || j != '8' || a != '9' || b != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}

				}
			}
			// }
			// }
			// }
		}
	}
	z01.PrintRune('\n')

}

// z01.PrintRune('7')
// z01.PrintRune('8')
// z01.PrintRune('9')
// z01.PrintRune('\n')

// }
// }
