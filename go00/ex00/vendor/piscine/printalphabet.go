package piscine

import (
	"ft"
)

func PrintAlphabet() error {
	for c := 'a'; c <= 'z'; c++ {
		if err := ft.PrintRune(c); err != nil {
			return err
		}
	}
	if err := ft.PrintRune('\n'); err != nil {
		return err
	}
	return nil
}
