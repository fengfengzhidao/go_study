package main

import "fmt"

func main() {
	// 声明一个string信道
	var ch chan string = make(chan string, 2)

	ch <- "枫枫" // 写入数据到信道中
	ch <- "知道"

	s := <-ch // 从信道读取数据
	fmt.Println(s)
	ss, ok := <-ch
	fmt.Println(ss, ok)

	close(ch)

}
