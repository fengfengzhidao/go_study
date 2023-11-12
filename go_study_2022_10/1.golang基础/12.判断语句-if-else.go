package main

import "fmt"

func main() {

	fmt.Println("请输入你的年龄：")
	var age uint8
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你已经成年了")
	} else if age >= 12 {
		fmt.Println("大于12小于18")
	} else {
		fmt.Println("未成年小朋友！")
	}

	// if简短语句
	if i := 3; i > 0 {
		fmt.Println("i > 0")
	}

}
