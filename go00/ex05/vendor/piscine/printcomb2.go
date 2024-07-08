package piscine

import "ft"

func printNum(num int) {
	ft.PrintRune(rune(num/10 + '0'))
	ft.PrintRune(rune(num%10 + '0'))
}

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printNum(i)
			ft.PrintRune(' ')
			printNum(j)
			if i != 98 || j != 99 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
