package main

import (
	"fmt"

	piscine "./factorial"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
}
