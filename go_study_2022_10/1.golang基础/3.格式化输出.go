package main

import "fmt"

func main() {

	// fmt.Print
	fmt.Print("xx")
	fmt.Print(123)

	// fmt.Println
	fmt.Println("xx")
	fmt.Println(123)

	// fmt.Printf
	// 你的名字是： ，今年年龄是
	fmt.Println("你的名字是：" + "张三" + "，今年年龄是：" + "21")
	fmt.Printf("你的名字是：%s，今年年龄是：%d\n", "张三", 21)

	fmt.Printf("二进制：%b, 八进制：%o %O, 十进制：%d, 十六进制：%x %X\n", 10, 10, 10, 10, 10, 10)

	fmt.Printf("%.2f\n", 3.1415)
	fmt.Printf("%T\n", 3.1415)
	fmt.Printf("%v\n", 3.1415)

}
