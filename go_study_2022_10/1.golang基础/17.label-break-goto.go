package main

import "fmt"

func main() {
	//start:
	//	for i := 1; i <= 5; i++ {
	//		for j := 1; j <= 6; j++ {
	//			fmt.Printf("* ")
	//			if i == 3 && j == 2 {
	//				break start
	//			}
	//		}
	//		fmt.Println()
	//	}

	//	for i := 1; i <= 5; i++ {
	//		for j := 1; j <= 6; j++ {
	//			fmt.Printf("* ")
	//			if i == 3 && j == 2 {
	//				goto start
	//			}
	//		}
	//		fmt.Println()
	//	}
	//start:

	//	var name string
	//start:
	//	fmt.Println("请输入用户名：")
	//	fmt.Scanln(&name)
	//	if name != "fengfeng" {
	//		fmt.Println("用户名错误！")
	//		goto start
	//	}
	//	fmt.Println("登录成功")

	fmt.Println(1)
	goto end
	fmt.Println(2)
	fmt.Println(3)
end:
	fmt.Println(4)

}
