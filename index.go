package piscine

func Index(s string, toFind string) int {
	sS := []rune(s)
	sF := []rune(toFind)
	kS := 0
	kF := 0
	for index := range sF {
		index = index
		kF++
	}
	if kF == 0 {
		return 0
	}
	for index := range sS {
		index = index
		kS++
	}
	for index, letter := range sS {
		if letter == sF[0] && kS >= kF+index-1 {
			a := 1
			for b := 1; b < kF; b++ {
				if sF[b] == sS[index+b] {
					a++
				}
			}
			if a == kF {
				return index
			}
		}
	}
	return -1
}
