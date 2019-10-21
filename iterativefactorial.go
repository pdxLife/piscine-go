package piscine

var arg int

func IterativeFactorial(arg int) int {

	result := 1

	for i := 1; i <= arg; i++ {
		result = result * i
	}
	return result
}
