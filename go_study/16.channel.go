package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan <- money

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
	go pay("张三", 2, &wait)
	go pay("王五", 3, &wait)
	go pay("李四", 5, &wait)

	go func() {
		defer close(moneyChan)
		wait.Wait()
	}()

	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

	//time.Sleep(2 * time.Second)

	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
}
