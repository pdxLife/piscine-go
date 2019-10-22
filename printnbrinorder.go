package piscine

import (
	"github.com/01-edu/z01"
)

func Digits(i int) (digit []int) {
	for i > 0 {
		if i == 0 {
			digit = append(digit, 0)
		} else {
			digit = append(digit, i%10)
		}
		i /= 10
	}
	return digit
}

func Sort(table []int) []int {
	int := 0
	for _, a := range table {
		if a == a {
			int++
		}
	}
	for a := 0; a < int; a++ {
		for b := 0; b < int; b++ {
			if table[b] > table[a] {
				table[a] = table[a] + table[b]
				table[b] = table[a] - table[b]
				table[a] = table[a] - table[b]
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for _, c := range Sort(Digits(n)) {
			z01.PrintRune(rune(c) + '0')
		}
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
