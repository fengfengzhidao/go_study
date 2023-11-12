package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
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
	wait.Add(4)
	// 主线程结束，协程函数跟着结束
	go shopping("张三", &wait)
	go shopping("王五", &wait)
	go shopping("李四", &wait)
	go shopping("王五", &wait)

	wait.Wait()
	//time.Sleep(2 * time.Second)

	fmt.Println("购买完成", time.Since(startTime))
}
