package main

import "fmt"

func main() {
	fmt.Println("A" > "a")      // 字符串就是比ascii值
	fmt.Println(2 > 1 || 1 > 2) // 有true为true
	fmt.Println(2 > 1 && 1 > 2) // 有false为false
	fmt.Println(!(2 > 1))

	// 逻辑短路
	// && 前面如果是false，后面就不判断了，直接为false
	a, b := 10, 0
	if a < 0 && a/b > 0 {
		fmt.Println("ok")
	}

	// || 如果前面是true，后面也不判断了，直接为true
	if a > 0 || a/b > 0 {
		fmt.Println("ok")
	}

}
