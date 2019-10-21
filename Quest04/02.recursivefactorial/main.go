package main

import (
	"fmt"

	piscine "./recursivefactorial"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
