package piscine

func StrLen(str string) int {

	length := 0
	for index, _ := range str {
		length = index + 1
	}
	// fmt.Println(length)
	return length
}

// func StrLen(str string) int {
// 	length := 0
// 	a := []rune(str)
// 	for p, _ := range a {
// 		length = p + 1
// 	}
// 	return length
// }
