package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	// str := []rune(str)

	// fmt.Println(args[1:])
	// fmt.Println(str)

	// for _, letter := range str {
	// 	for i := 0; i < len(str); i++ {
	// 		if str[i] <= 'm' {
	// 			return rune(str[i-13])
	// 		} else {
	// 			return str[i+13]
	// 		}

	// 		fmt.Println(letter)
	// 	}
	// }
	// fmt.Println('\n')

	for i := 0; i < len(args); i++ {
		fmt.Println(os.Args[0])
	}
	// fmt.Println('\n')
}

func Rot13(str string) string {

	arg := []rune(str)
	for i := 0; i < len(arg); i++ {
		if arg[i] >= 'A' && arg[i] <= 'N' {
			arg[i] = arg[i] + 13
		} else if arg[i] >= 'N' && arg[i] <= 'Z' {
			arg[i] = arg[i] - 13
		} else if arg[i] >= 'a' && arg[i] <= 'n' {
			arg[i] = arg[i] + 13
		} else if arg[i] >= 'n' && arg[i] <= 'z' {
			arg[i] = arg[i] - 13
		}
	}
	return string(arg)
}
