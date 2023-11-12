package main

import "fmt"

func main() {
	var ch1 chan int = make(chan int, 2)
	var ch2 chan int = make(chan int, 2)
	var ch3 chan int = make(chan int, 2)
	ch1 <- 1
	ch2 <- 2
	ch3 <- 3

	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case v := <-ch3:
		fmt.Println(v)
	default:
		fmt.Println("没有数据")
	}

}
