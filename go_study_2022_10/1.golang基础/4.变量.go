package main

import (
	"fmt"
	"go_study/fengfengStudy/version"
)

// 批量声明
var (
	m78      string
	userName int = 10
	UserName bool
)

func main() {
	var name string
	var age int
	var score float32
	var isSuper bool
	// 针对基础类型 如果声明之后不赋值，会有默认值
	fmt.Println(name, age, score, isSuper)

	name = "fengfeng"
	fmt.Println(name)
	// 声明的同时赋值
	var title string = "枫枫知道"
	fmt.Println(title)

	// 自动推断类型
	var userName = "枫枫"
	fmt.Println(userName)

	// 声明赋值
	userAddr := "长沙市"
	fmt.Println(userAddr)

	fmt.Println(m78)

	fmt.Println(version.Version)
	fmt.Println(version.STATE)
}
