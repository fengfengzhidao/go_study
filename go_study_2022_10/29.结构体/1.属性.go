package main

import "fmt"

func main() {
	type User struct {
		Name     string
		Age      int
		password string
	}
	var u1 User = User{"张三", 21, "*7gd783"}
	fmt.Printf("u1:%#v, %T\n", u1, u1)
	var u2 User = User{
		//Age:      21,   // 如果不传就是默认值
		Name:     "张三",
		password: "*7gd783",
	}
	fmt.Printf("u2:%#v, %T\n", u2, u2)
	// 第三种
	var u3 User

	u3.Name = "zhangsan"
	u3.Age = 21
	u3.password = "#Xjg84k"
	fmt.Printf("u3:%#v, %T\n", u3, u3)
	u4 := User{"张三", 21, "*7gd783"}
	fmt.Printf("u4:%#v, %T\n", u4, u4)

	fmt.Println(u4.Name, u4.Age)
}
