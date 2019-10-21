package piscine

//var arg int

func RecursiveFactorial(nb int) int {

	if nb >= 0 && nb < 20 {

		if nb == 1 {
			return 1
		} else if nb > 1 {
			return nb * RecursiveFactorial(nb-1)
		} else {
			return 1
		}
	}
	return 0
}

// package piscine

// var arg int

// func IterativeFactorial(nb int) int {

// 	if nb >= 0 && nb < 25 {

// 		result := 1

// 		for i := 1; i <= nb; i++ {
// 			result = result * i
// 		}
// 		return result
// 	} else {
// 		return 0
// 	}
// }
