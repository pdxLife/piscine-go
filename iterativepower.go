package piscine

func IterativePower(nb int, power int) int {
	p := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := power; i > 1; i-- {
		p = nb * p
	}
	return p
}

// func IterativePower(nb int, power int) int {

// 	// if nb >= 0 && nb < 20 {

// 	if power == 0 {
// 		return 1
// 	}
// 	if power > 1 {
// 		result := 0
// 		for i := 0; i <= power; i++ {
// 			result = result * nb
// 		}
// 		return result
// 	}
// 	if power < 0 {
// 		return 0
// 	}
// 	return 1
// }
