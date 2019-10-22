package piscine

func ToUpper(s string) string {
	word := []byte(s)
	for index, letter := range word {
		if letter <= 122 && letter >= 97 {
			word[index] = letter - 32
		}
	}
	return string(word)
}
