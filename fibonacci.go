package piscine

func Fibonacci(index int) int {
	// var result = 0
	// for i := 0; i <= index; i++ {

	// 	// a = result
	// 	result = result + i
	// }
	// return result

	if index <= 1 {
		return -1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)

}
