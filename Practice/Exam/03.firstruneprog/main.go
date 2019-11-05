package main

import (
	"fmt"

	piscine "./firstruneprog"
)

func main() {
	str := "Hello World!"
	nb := piscine.FirstRune(str)
	fmt.Println(nb)
}
