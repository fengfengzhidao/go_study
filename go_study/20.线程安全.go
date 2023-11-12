package main

import (
	"fmt"
	"sync"
)

var sum1 int

var wait1 sync.WaitGroup
var lock sync.Mutex

func add1() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum1++
	}
	lock.Unlock()
	wait1.Done()
}
func sub1() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum1--
	}
	lock.Unlock()
	wait1.Done()
}

func main() {
	wait1.Add(2)
	go add1()
	go sub1()
	wait1.Wait()

	fmt.Println(sum1)

}
