package piscine

func RecursivePower(nb int, power int) int {

	if nb >= 0 && nb < 20 {

		if power == 0 {
			return 1
		} else if nb > 1 {
			return nb * RecursivePower(nb, power-1)
		} else {
			return 0
		}
	}
	return 0
}

// func RecursiveFactorial(nb int) int {

// 	if nb >= 0 && nb < 20 {

// 		if nb == 1 {
// 			return 1
// 		} else if nb > 1 {
// 			return nb * RecursiveFactorial(nb-1)
// 		} else {
// 			return 1
// 		}
// 	}
// 	return 0
// }
