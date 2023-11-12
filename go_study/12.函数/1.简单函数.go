package main

import "fmt"

func sayHello() {
	fmt.Println("你好")
}

func param1(id string) {
	fmt.Println(id)
}
func param2(id string, userName string) {
	fmt.Println(id)
}
func param3(id, userName string) {
	fmt.Println(id)
}
func add(numberList ...int) {
	var sum int
	for _, i2 := range numberList {
		sum += i2
	}
	fmt.Println(sum)
}

func r1() {
	// 没有返回值
	return
}

func r2() bool {
	var ok bool
	return ok
}

func r3() (string, bool) {
	if 1 > 2 {
		return "", false
	}
	if 1 > 2 {
		return "", false
	}
	return "", true
}
func r4() (val string, ok bool) {
	if 1 > 2 {
		val = "12"
		return
	}
	return
}

func main() {
	//sayHello()
	//param1("123")

	//add(1, 2)
	//add(1, 2, 3)
	//add(1, 2, 3, 4)

	//var getName = func() string {
	//	return "枫枫"
	//}
	//var setName = func(name string) {
	//	fmt.Println(name)
	//	return
	//}
	//setName("zhangsan")
	//fmt.Println(getName())

	fmt.Println("请输入你要执行的操作：")
	fmt.Println("1 登录")
	fmt.Println("2 注册")
	fmt.Println("3 个人中心")
	var index int
	fmt.Scan(&index)

	var funMap = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	}

	//switch index {
	//case 1:
	//	login()
	//case 2:
	//	register()
	//case 3:
	//	userCenter()
	//}
}

func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")
}
