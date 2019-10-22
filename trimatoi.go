package piscine

func FirstCheck(s string) bool {
	abc := false
	for _, a := range s {
		if a >= '0' && a <= '9' {
			abc = true
			break
		}
	}
	return abc
}

func SecondCheck(s string) bool {
	neg := false
	for _, a := range s {
		if a >= '0' && a <= '9' {
			break
		}
		if a == '-' {
			neg = true
		}
	}
	return neg
}

func TrimAtoi(s string) int {
	x := 0
	if FirstCheck(s) == true {
		for _, a := range s {
			if a >= '0' && a <= '9' {
				b := 0
				for i := '1'; i <= a; i++ {
					b++
				}
				x = x*10 + b
			}
		}
	}
	if SecondCheck(s) {
		x *= -1
	}
	return x
}
