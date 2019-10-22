package piscine

func LastRune(s string) rune {
	str := []rune(s)
	a := 0
	for _, b := range str {
		if b == b {
		}
		a++
	}
	return str[a-1]
}
