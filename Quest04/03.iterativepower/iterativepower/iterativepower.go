package piscine

func IterativePower(nb int, power int) int {

	if nb >= 0 && nb < 20 {

		if power == 0 {
			return 1
		} else if power > 1 {
			result := 0
			for i := 0; i <= power; i++ {
				result = result * i
			}
			return result
		} else {
			return 1
		}
	}
	return 0
}

// if nb >= 0 && nb < 25 {

// 	result := 1

// 	for i := 1; i <= nb; i++ {
// 		result = result * i
// 	}
// 	return result
// } else {
// 	return 0
