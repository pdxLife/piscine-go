package piscine

var arg int

func RecursiveFactorial(arg int) int {

	if arg == 1 {
		return 1
	} else if arg > 1 {
		return arg * RecursiveFactorial(arg-1)
	} else {
		// result := 1

		// for i := 1; i <= arg; i++ {
		// 	result = result * i
		// }
		return 0
	}
}
