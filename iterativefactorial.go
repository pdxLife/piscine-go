package piscine

var arg int

func IterativeFactorial(nb int) int {

	if nb > 0 && nb < 25 {

		result := 1

		for i := 2; i <= nb; i++ {
			result = result * i
		}
		return result
	} else {
		return 0
	}
}
