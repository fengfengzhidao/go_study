package main

import (
	"fmt"
	"runtime/debug"
)

func read() {

	var list []int = []int{1, 2}
	fmt.Println(list[2])
}

func main1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	read()
}

func main() {
	main1()
	// 正常逻辑
	fmt.Println("正常逻辑")
}
