package main

import (
	"fmt"

	piscine ".."
)

func main() {
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
