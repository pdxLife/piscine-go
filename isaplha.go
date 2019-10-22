package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)
	for index, i := range runes {
		if !((i >= '0' && i <= '9') || (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || i == ' ') {
			index++
			index--
			return false
		}
	}
	return true
}
