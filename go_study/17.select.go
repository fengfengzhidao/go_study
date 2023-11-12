package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int)    // 声明并初始化一个长度为0的信道
var nameChan1 = make(chan string)  // 声明并初始化一个长度为0的信道
var doneChan = make(chan struct{}) // 声明一个用于关闭的信道

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

// 协程
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	// 现在的模式，就是购物接力
	//shopping("张三")
	//shopping("王五")
	//shopping("李四")
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
	go send("张三", 2, &wait)
	go send("王五", 3, &wait)
	go send("李四", 5, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)

}


