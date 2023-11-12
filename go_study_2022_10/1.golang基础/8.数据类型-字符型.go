package main

import "fmt"

func main() {
	var c0 uint8 = 65
	var c1 byte = 65
	var c2 = 'a'
	var c3 rune = '中'

	fmt.Printf("c0 的码值：%v, 字符为：%c, 类型是%T。\n", c0, c0, c0)
	fmt.Println(c0 == c1) // true
	fmt.Printf("c1 的码值：%v, 字符为：%c, 类型是%T。\n", c1, c1, c1)
	fmt.Printf("c2 的码值：%v, 字符为：%c, 类型是%T。\n", c2, c2, c2)
	fmt.Printf("c3 的码值：%v, 字符为：%c, 类型是%T。\n", c3, c3, c3)

}
