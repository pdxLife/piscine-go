package piscine

func Capitalize(s string) string {
	sAsRune := []rune(s)
	for index, letter := range sAsRune {
		if checkAlphNum(letter) {
			if index == 0 || checkAlphNum(sAsRune[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					sAsRune[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					sAsRune[index] = letter + 32
				}
			}
		}
	}
	return string(sAsRune)
}

func checkAlphNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
