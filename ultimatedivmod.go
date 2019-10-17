package piscine

var c, d int

func UltimateDivMod(a *int, b *int) {
	c = *a / *b
	d = *a % *b
	*a = c
	//c = a
	*b = d
	//d = b
}
