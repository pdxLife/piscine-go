package main

import (
	"fmt"

	piscine "./strlenprog"
)

func main() {
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
