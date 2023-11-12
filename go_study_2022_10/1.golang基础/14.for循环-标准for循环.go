package main

import "fmt"

func main() {

	for i := 0; i < 10; i += 2 {
		fmt.Println("打印", i)
	}
	i := 0
	for i < 5 {
		fmt.Println("for ", i)
		i++
	}

	for {
		fmt.Println("xx ", i)
		i--
		if i == 0 {
			// 退出循环
			break

		}

	}

}
