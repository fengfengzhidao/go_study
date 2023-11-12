package main

import "fmt"

func main() {
	var ch chan int = make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	for i2 := range ch {
		fmt.Println(i2)
	}
	/*
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/

}
