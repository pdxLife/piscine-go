package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 {
			z01.PrintRune(111) //print 0
			z01.PrintRune(10)  //print new line
		}
		if x > 1 {
			z01.PrintRune(111)
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune(45) //print -
			}
			z01.PrintRune(111)
			z01.PrintRune(10)
		}
		if x == 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune(124) //print |
				z01.PrintRune(10)
			}
		}
		if y >= 1 && x != 1 {
			for yl := 0; yl < y-2; yl++ {
				z01.PrintRune(124)
				for xl2 := 0; xl2 < x-2; xl2++ {
					z01.PrintRune(32) //print space
				}
				z01.PrintRune(124)
				z01.PrintRune(10)
			}
		}
		if x == 1 && y != 1 {
			z01.PrintRune(111)
			z01.PrintRune(10)
		}
		if x > 1 && y != 1 {
			z01.PrintRune(111)
			for xl := 0; xl < x-2; xl++ {
				z01.PrintRune(45)
			}
			z01.PrintRune(111)
			z01.PrintRune(10)
		}
	}
}

// func Raid1a(x, y int) {
// 	for i := 0; i <= x; i++ {
// 		for j := 0; j <= y; j++ {
// 			if i == 1 || i == y {
// 				if j == 1 || j == x {
// 					z01.PrintRune('o') //Print o Ascii 111
// 				} else {
// 					z01.PrintRune('-') //Print - Ascii 45
// 				}
// 			} else {
// 				if j == 1 || j == x {
// 					z01.PrintRune(' ') //Print Space Ascii 32
// 				} else {
// 					z01.PrintRune('|') //Print | Ascii 124
// 				}
// 			}
// 		}

// 	}
// 	z01.PrintRune(10)
// }
