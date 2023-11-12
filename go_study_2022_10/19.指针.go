package main

import "fmt"

func setName(name *string) {
	// 字符串变量name映射的地址
	*name = "知道"
	fmt.Printf("函数内部：%p\n", name)
}

func main() {
	name := "枫枫" // 值
	fmt.Printf("函数外部：%p\n", &name)
	setName(&name)
	fmt.Println(name)
}
