package main

import "fmt"

func main() {
	var s1 string = `枫枫\n知道`
	var s2 string = "枫枫\n知道"

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1)) // 汉字三字节

	fmt.Println(true)
	fmt.Println(false)

}
