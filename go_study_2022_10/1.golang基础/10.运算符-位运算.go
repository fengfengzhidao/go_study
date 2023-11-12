package main

import "fmt"

func main() {
	var b uint8 = 0b11001100
	// 右移两位
	fmt.Printf("b>>2: %b\n", b>>2) // 0b 00110011
	// 左移
	fmt.Printf("b<<2: %b\n", b<<2) // 0b 00110000

	var (
		c1 uint8 = 0b10101010
		c2 uint8 = 0b11001100
	)
	// 按位与 全1为1
	fmt.Printf("c1&c2: %b\n", c1&c2) // 10001000
	// 按位或 有1为1
	fmt.Printf("c1|c2: %b\n", c1|c2) // 11101110
	// 按位异或 相同为0 不同为1
	fmt.Printf("c1^c2: %b\n", c1^c2) // 01100110

	fmt.Println(4 >> 2) // 00100
	fmt.Println(2 << 2) // 00010  01000
	fmt.Println(10 | 5) //
}
