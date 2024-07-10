package piscine

/*
Write a function that does the following:
• This function will divide an int a by another int b.
• The result of this division will be stored in the int pointed by div.
• The remainder of this division will be stored in the int pointed by mod.
• Expected function
*/
func DivMod(a int, b int, div *int, mod *int) {

	if b == 0 {
		return
	}
	*div = a / b
	*mod = a % b
}
