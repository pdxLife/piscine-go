package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	for index, i := range runes {
		if !((i >= 'a' && i <= 'z') || i == ' ') {
			index = index + 0
			return false
		}
	}
	return true
}
