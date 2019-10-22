package piscine

func ToLower(s string) string {
	word := []byte(s)
	for index, letter := range word {
		if letter <= 90 && letter >= 65 {
			word[index] = letter + 32
		}
	}
	return string(word)
}
