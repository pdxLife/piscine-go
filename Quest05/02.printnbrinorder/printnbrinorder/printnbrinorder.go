package piscine

import "github.com/01-edu/z01"

func IntToDigit(n int) (digits []int) {

	for n > 0 {
		if n == 0 {
			digits = append(digits, n)
		} else {
			digits = append(digits, n%10)
		}
		n = n / 10
	}
	return

}

func sort(array []int) []int {

	counter := 0
	for range array {

		counter = counter + 1
	}
	for i := counter; i > 0; i-- {
		for j := 1; j < 1; j++ {

			if array[j] < array[j+1] {
				value := array[j]
				array[j] = array[j-1]
				array[j-1] = value
			}
		}
	}
	return array
}

func PrintNbrInOrder(n int) {

	if n == 0 {

		z01.PrintRune('0')
		return
	}

	for _, c := range sort(IntToDigit(n)) {

		z01.PrintRune(rune(c) + '0')
	}

}

// counter := 0

// // for index, letter := range str {

// for _, letter := range str {

// 	if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
// 		counter++

// 		//fmt.Printf("index is %v     letter is &w \n", index, letter)
// 	}
// }
// //fmt.Printf("Counter value is %v \n", counter)
// //fmt.Println()
// return counter
