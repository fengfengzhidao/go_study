package main

import "fmt"

func main() {
	fmt.Print("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Print("请输入你的年龄：")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
