package piscine

func Fibonacci(index int) int {
	// var result = 0
	// for i := 0; i <= index; i++ {

	// 	// a = result
	// 	result = result + i
	// }
	// return result

	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
