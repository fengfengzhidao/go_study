package main

import (
	"fmt"
	"sort"
)

func main() {
	//var nameList []string
	nameList := make([]string, 0)
	//nameList = append(nameList, "枫枫")
	//nameList = append(nameList, "张三")
	fmt.Println(nameList == nil)

	ageList := make([]int, 3)
	fmt.Println(ageList)

	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])
	fmt.Println(array[1:2])

	var ints = []int{4, 8, 3, 2}
	sort.Ints(ints) // 升序
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // 降序
	fmt.Println(ints)

}
