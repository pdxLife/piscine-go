package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	ok := true
	for _, i := range runes {
		if i >= 0 && i <= 31 {
			ok = false
			return ok
		}
	}
	return ok
}

//package piscine

//func IsPrintable(str string) bool {
//	runes := []rune(str)
//	for index, i := range runes {
//		if !(i >= ' ') {
//			index = index + 0
//			return false
//		}
//	}
//	return true
//}
