package piscine

func FindNextPrime(nb int) int {

	result := nb
	// if nb < 2 {
	// 	return
	// }
	for i := 2; i*i <= nb; i++ {

		if nb%i == 0 {
			return nb
		}
	}
	return result
}
