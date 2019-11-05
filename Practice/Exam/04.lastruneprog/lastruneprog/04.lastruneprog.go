package piscine

func LastRune(s string) rune {

	a := []rune(s)
	length := 0
	for index, _ := range s {
		length = index + 1
	}
	return a[length-1]
}
