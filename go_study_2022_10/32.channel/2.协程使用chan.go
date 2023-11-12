package main

import (
	"fmt"
	"sync"
)

var ch chan int = make(chan int, 10)
var wg = sync.WaitGroup{}

func pushNum() {

	for i := 0; i < 5; i++ {
		ch <- i
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go pushNum()
	go pushNum()
	wg.Wait()
	close(ch)
	for {
		res, ok := <-ch
		if !ok {
			break

		}
		fmt.Println(res)
	}

}
