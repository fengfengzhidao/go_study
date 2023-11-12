package main

import "fmt"

func main() {
	// 1
	// 2
	// 3
	// 4
	// 100
	//var sum int
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	// 死循环
	//for i := 1; true; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(2 * time.Second)
	//}

	//for true {
	//	fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	//}
	//for {
	//	fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	//}

	//var sum int
	//var i int = 1
	//for i <= 100 {
	//	sum += i
	//	i++
	//}
	//fmt.Println(sum)

	//var sum int
	//var i int = 1
	//for {
	//	sum += i
	//	i++
	//	if i == 101 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	//var list = []string{"枫枫", "知道"}
	//for i := 0; i < len(list); i++ {
	//	fmt.Println(i, list[i])
	//}
	//for index, item := range list {
	//	fmt.Println(index, item)
	//}

	//var userMap = map[string]string{"name": "枫枫"}
	//for key, value := range userMap {
	//	fmt.Println(key, value)
	//}

	// 1*1=1
	// 2*1=2  2*2=4
	// 9*1=9 ... 9*9=81

	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d\t", i, j, i*j)
	//	}
	//	fmt.Println()
	//}

	for i := 0; i <= 10; i++ {

		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
