package main

import (
	"fmt"

	piscine "./trimatoi"
)

func main() {
	s := "12345"
	s2 := "str123ing45"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "sd+x1fa2W3s4"
	s6 := "sd-x1fa2W3s4"
	s7 := "sdx1-fa2W3s4"

	n := piscine.TrimAtoi(s)
	n2 := piscine.TrimAtoi(s2)
	n3 := piscine.TrimAtoi(s3)
	n4 := piscine.TrimAtoi(s4)
	n5 := piscine.TrimAtoi(s5)
	n6 := piscine.TrimAtoi(s6)
	n7 := piscine.TrimAtoi(s7)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}
