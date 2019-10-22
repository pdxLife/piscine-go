// package piscine

// func FindNextPrime(nb int) int {

// 	result := nb
// 	// if nb < 2 {
// 	// 	return
// 	// }
// 	for i := 2; i*i <= nb; i++ {

// 		if nb%i == 0 {
// 			return nb
// 		}
// 	}
// 	return result
// }

package piscine

func check(a int) bool {
	if a <= 1 {
		return false
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for {
		if check(nb) {
			return nb
		}
		nb++
	}
}
