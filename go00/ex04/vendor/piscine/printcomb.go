package piscine

import "ft"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				ft.PrintRune(rune(i + '0'))
				ft.PrintRune(rune(j + '0'))
				ft.PrintRune(rune(k + '0'))
				if i != 7 || j != 8 || k != 9 {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
