package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))                // 12345
	fmt.Println(piscine.BasicAtoi("0000000012345"))        // 12345
	fmt.Println(piscine.BasicAtoi("000000"))               // 0
	fmt.Println(piscine.BasicAtoi("0"))                    // 0
	fmt.Println(piscine.BasicAtoi("1"))                    // 1
	fmt.Println(piscine.BasicAtoi("9876543210"))           // 0
	fmt.Println(piscine.BasicAtoi("42"))                   // 42
	fmt.Println(piscine.BasicAtoi("9223372036854775807"))  // 9223372036854775807
	fmt.Println(piscine.BasicAtoi("9223372036854775808"))  // 0 (overflow)
	fmt.Println(piscine.BasicAtoi("18446744073709551615")) // 0 (overflow)
	fmt.Println(piscine.BasicAtoi("123abc"))               // 0 (invalid characters)
	fmt.Println(piscine.BasicAtoi("abc123"))               // 0 (invalid characters)
	fmt.Println(piscine.BasicAtoi(""))                     // 0 (empty string)
}
