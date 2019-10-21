package student

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 {
			z01.PrintRune('/')  //ascii 47
			z01.PrintRune('\n') //ascii 10
		}
		if x > 1 {
			z01.PrintRune('/')
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune('*') //ascii 42
			}
			z01.PrintRune(92)   //ascii 92
			z01.PrintRune('\n') //ascii 10
		}
		if x == 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune('*') //
				z01.PrintRune('\n')
			}
		}
		if y >= 1 && x != 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune('*')
				for xl2 := 0; xl2 < x-2; xl2++ {
					z01.PrintRune(32) //ascii 32
				}
				z01.PrintRune(42)
				z01.PrintRune(10)
			}
		}
		if x == 1 && y != 1 {
			z01.PrintRune(92)
			z01.PrintRune(10)
		}
		if x > 1 && y != 1 {
			z01.PrintRune(92)
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune(42)
			}
			z01.PrintRune(47)
			z01.PrintRune(10)
		}
	}
}
