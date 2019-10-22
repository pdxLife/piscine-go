package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	for a, b := range s1 {
		if a == n-1 {
			return b
		}
	}
	return 0
}
