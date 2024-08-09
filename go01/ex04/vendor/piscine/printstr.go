package piscine

import "ft"

func PrintStr(s string) {
	for _, r := range s {
		if err := ft.PrintRune(r); err != nil {
			return
		}
	}
	ft.PrintRune('\n')

}
