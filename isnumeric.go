package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	for index, i := range runes {
		if !((i >= '0' && i <= '9') || i == ' ') {
			index = index + 0
			return false
		}
	}
	return true
}
