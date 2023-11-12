package main

import (
	"fmt"
	"math"
)

func main() {
	var i8 int8
	i8 = 127

	// 0 0 0 0 0 0 0 0      0
	// 0 0 0 0 0 0 0 1      1   2 ** 0
	// 0 0 0 0 0 0 1 0      2   2 ** 1
	// 0 0 0 0 0 0 1 1      3
	// 0 0 0 0 0 1 0 0      4   2 ** 2
	// 0 0 0 0 1 0 0 0      4   2 ** 3
	// 0 0 0 1 0 0 0 0      4   2 ** 4
	// 0 0 1 0 0 0 0 0      4   2 ** 5
	// 0 1 0 0 0 0 0 0      4   2 ** 6
	// 0 1 1 1 1 1 1 1      4   2 ** 7 - 1

	fmt.Println(math.Pow(2, 15))
	fmt.Println(i8)

	var b uint8 = 0b00100000
	fmt.Printf("0b%b, %d", b, b)

	var h uint8 = 0xa1 // 10 * 16 + 1 * 1
	fmt.Printf("0x%x, %d", h, h)

}
