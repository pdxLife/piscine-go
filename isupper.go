package piscine

func IsUpper(str string) bool {
	runes := []rune(str)
	for index, i := range runes {
		if !((i >= 'A' && i <= 'Z') || i == ' ') {
			index = index + 0
			return false
		}
	}
	return true
}
