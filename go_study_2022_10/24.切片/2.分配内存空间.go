package main

import "fmt"

func main() {
	var s []int
	// 1.使用make  创建一个长度为3的int切片，默认值是0
	s = make([]int, 3, 4)
	fmt.Println(s, len(s), cap(s))
	// 2.系统自动创建底层数组
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1, len(s1), cap(s1))
}
