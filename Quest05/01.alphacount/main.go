package main

import (
	"fmt"

	piscine "./alphacount"
)

func main() {
	str := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(str)
	fmt.Println(nb)
}
