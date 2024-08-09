package piscine

func UltimateDivMod(a *int, b *int) {

	if a == nil || b == nil || *b == 0 {
		return
	}
	tmpDiv := *a / *b
	*b = *a % *b
	*a = tmpDiv
}
