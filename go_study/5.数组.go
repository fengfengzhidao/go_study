package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"枫枫", "张三", "李四"}
	fmt.Println(nameList)
	// "枫枫"   "张三"   "李四"
	//   0       1        2
	//  -3      -2       -1
	fmt.Println(nameList[0])
	fmt.Println(nameList[2])
	fmt.Println(nameList[len(nameList)-1])

	nameList[0] = "枫枫知道"
	fmt.Println(nameList)
	//fmt.Println(nameList[-1]) go语言不支持

}
