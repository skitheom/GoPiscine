package piscine

import "ft"

func printDigits(comb []rune, numOfDigits int) {
	for i := 0; i < numOfDigits; i++ {
		ft.PrintRune(comb[i])
	}
	if comb[0] != rune(10-numOfDigits+'0') {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func computeComb(comb []rune, position int, numOfDigits int) {

	if position == numOfDigits {
		printDigits(comb, numOfDigits)
		return
	}
	start := '0'
	if position > 0 {
		start = comb[position-1] + 1
	}
	for i := start; i <= '9'; i++ {
		comb[position] = rune(i)
		computeComb(comb, position+1, numOfDigits)
	}
}

func PrintCombN(n int) {

	if n < 1 || n > 9 {
		return
	}
	var comb [10]rune
	computeComb(comb[:], 0, n)
	ft.PrintRune('\n')
}
