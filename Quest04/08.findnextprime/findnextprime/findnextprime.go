package piscine

func IsPrime(nb int) bool {

	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {

	for i := nb; i++{
		if IsPrime(i) == true {
			return i
		}
	}
}

// package piscine

// func check(a int) bool {
// 	if a <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= a; i++ {
// 		if a%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func FindNextPrime(nb int) int {
// 	for {
// 		if check(nb) {
// 			return nb
// 		}
// 		nb++
// 	}
// }
