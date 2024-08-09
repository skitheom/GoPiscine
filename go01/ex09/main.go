package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi2("12345"))         // 12345
	fmt.Println(piscine.BasicAtoi2("0000000012345")) // 12345
	fmt.Println(piscine.BasicAtoi2("012 345"))       // 0
	fmt.Println(piscine.BasicAtoi2("Hello World!"))  // 0
	fmt.Println(piscine.BasicAtoi2("42"))            // 42
}
