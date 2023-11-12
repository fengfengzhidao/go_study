package main

import "fmt"

func main() {

	// switch case的第一种用法
	//
	//fmt.Println("请输入你的年龄：")
	//var age uint8
	//fmt.Scanln(&age)
	//
	//switch {
	//case age >= 18:
	//	fmt.Println("你已经成年了")
	//	fallthrough // 穿透，继续往下匹配
	//case age >= 12:
	//	fmt.Println("大于12小于18")
	//	break // 退出
	//default:
	//	fmt.Println("未成年小朋友！")
	//}
	//

	//switch case 的第二种用法

	fmt.Println("请输入星期(数字)：")
	var weekDay uint8
	fmt.Scanln(&weekDay)

	switch weekDay {
	case 1:
		fmt.Println("蛋炒饭")
	case 2:
		fmt.Println("海鲜大餐")
	case 3:
		fmt.Println("辣椒炒肉")
	default:
		fmt.Println("西北风")
	}

}
