package main

import (
	"fmt"
)

//1. 要求用户从命令行输入用户名，密码，确认密码
//2. 如果输入不为空且两次密码相同，则打印注册成功并结束程序，否则根据情况提示用户输入不能为空或两次密码不一致，并且要求用户重新输入

func main() {

	var (
		username   string
		password   string
		rePassword string
	)
	for {
		fmt.Println("请输入用户名:")
		fmt.Scanln(&username)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)
		fmt.Println("请再次输入密码:")
		fmt.Scanln(&rePassword)

		//if !(username != "" && password != "" && rePassword != "") {
		//	fmt.Println("输入不能为空！")
		//	continue
		//}

		if username == "" || password == "" || rePassword == "" {
			fmt.Println("输入不能为空！")
			continue
		}

		if password != rePassword {
			fmt.Println("两次密码不一致！")
			continue
		}
		fmt.Println("注册成功！")
		//os.Exit(0)
		break
	}

}
