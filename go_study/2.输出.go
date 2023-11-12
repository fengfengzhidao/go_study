package main

import "fmt"

func main() {
	fmt.Println("hello", "你好")
	fmt.Println("你好", "枫枫")
	fmt.Printf("%s 哇，你好美！\n", "枫枫")
	fmt.Printf("%d\n", 3)
	fmt.Printf("%.2f\n", 3.1415892)
	fmt.Printf("%T %T\n", "你好", 2.5)
	fmt.Printf("%v\n", "")
	fmt.Printf("%#v\n", "")

	var f = fmt.Sprintf("%.2f", 3.1415892)
	fmt.Println(f)
}
