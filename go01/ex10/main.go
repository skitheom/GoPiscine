package main

import (
	"fmt"
	"piscine"
)

func main() {
	// 通常のケース
	fmt.Println(piscine.Atoi("12345"))         // 12345
	fmt.Println(piscine.Atoi("0000000012345")) // 12345
	fmt.Println(piscine.Atoi("012 345"))       // 0
	fmt.Println(piscine.Atoi("Hello World!"))  // 0
	fmt.Println(piscine.Atoi("+1234"))         // 1234
	fmt.Println(piscine.Atoi("-1234"))         // -1234
	fmt.Println(piscine.Atoi("++1234"))        // 0
	fmt.Println(piscine.Atoi("--1234"))        // 0

	// 境界ケース
	fmt.Println(piscine.Atoi("9223372036854775807"))  // INT_MAX: 9223372036854775807
	fmt.Println(piscine.Atoi("-9223372036854775808")) // INT_MIN: -9223372036854775808
	fmt.Println(piscine.Atoi("9223372036854775808"))  // INT_MAX + 1: 0 (overflow)
	fmt.Println(piscine.Atoi("-9223372036854775809")) // INT_MIN - 1: 0 (overflow)
}
