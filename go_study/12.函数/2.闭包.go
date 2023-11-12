package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitSecond int) func(...int) int {
	//time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberList ...int) (sum int) {
		time.Sleep(time.Duration(awaitSecond) * time.Second)
		for _, i2 := range numberList {
			sum += i2
		}
		return sum
	}
}

func Copy(name string) {
	fmt.Printf("infun %p\n", &name)
	name = "枫枫不知道"
}

func Set(name *string) {
	fmt.Printf("infun %p\n", name)
	*name = "枫枫不知道"
	// 通过内存地址去找值
}

func main() {
	//add := awaitAdd(2)
	//t1 := time.Now()
	//sum := add(1, 2, 3)
	//subTime := time.Since(t1)
	//fmt.Println(sum, subTime)

	var name = "枫枫"
	fmt.Printf("%p\n", &name)
	Set(&name)
	fmt.Println(name)
}
