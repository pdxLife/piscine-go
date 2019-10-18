package piscine

func StrRev(s string) string {
	var k string
	i = 0
	for range s {
		i++
	}
	for j := range s {
		k = string(s[i]) + k
	}
	return k
}

//func reverse
