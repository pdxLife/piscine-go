package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := os.Args

	// arg2 := os.Args[i]

	length := 0

	for range arg {
		length++
	}

	// for i := 0; i < length; i++ {
	// for _, letter := range arg[i] {
	// 	z01.PrintRune(letter)
	// }
	z01.PrintRune(arg[length-1])
	z01.PrintRune('\n')

	// }

}
