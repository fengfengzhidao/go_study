package main

import (
	"fmt"
	"go_study/go_study/version"
)

func hello() {
	//fmt.Println(age)
	fmt.Println("你好")
}

var age = 12
var (
	s1 string = "1"
	s2 string = "2"
)

const Version string = "2.0.1"

func main() {
	fmt.Println(version.Version)
	//fmt.Println(version.fengfeng)
	// 先声明
	var name string
	// 再赋值
	name = "枫枫"
	fmt.Println(name)

	// 声明并赋值
	var name1 string = "枫枫"
	fmt.Println(name1)
	// 省略类型
	var name2 = "枫枫"
	fmt.Println(name2)

	// 声明并赋值 短声明符号
	name3 := "枫枫"
	fmt.Println(name3)

	hello()

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

}
