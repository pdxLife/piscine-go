
package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for index, i := range runes {
		if !(i >= ' ') {
			index = index + 0
			return false
		}
	}
	return true
}
