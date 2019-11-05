package main

import (
	"fmt"

	piscine "./lastruneprog"
)

func main() {
	str := "Hello World!"
	nb := piscine.LastRune(str)
	fmt.Println(nb)
}
