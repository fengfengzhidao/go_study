package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 2)
	// 可读chan
	var readCh <-chan int = ch
	// 可写chan
	var writeCh chan<- int = ch

	writeCh <- 1
	writeCh <- 2

	fmt.Println(<-readCh)
	fmt.Println(<-readCh)

}
