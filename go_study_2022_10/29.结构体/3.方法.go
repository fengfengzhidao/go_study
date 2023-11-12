package main

import "fmt"

//type User struct {
//	Name     string
//	Age      int
//	password string
//}
//
//func (u User) PrintName() string {
//	u.Name = "张三"
//	fmt.Println("printName方法： ", u)
//	fmt.Printf("printName方法内部：%p \n", &u)
//	return u.Name
//}

func main() {
	type User struct {
		Name     string
		Age      int
		password string

		PrintName func(u User) string
	}

	// 继承也会把方法继承到
	type Account struct {
		User
		Money float32
	}

	user := User{"王二麻子", 23, "xsdfe23", func(u User) string {
		u.Name = "张三"
		fmt.Println("printName方法： ", u)
		fmt.Printf("printName方法内部：%p \n", &u)
		return u.Name
	}}
	fmt.Printf("main: %p\n", &user)
	name := user.PrintName(user) // 是值传递，
	fmt.Println(name, user)

	account := Account{
		Money: 234,
		User:  user,
	}
	name = account.PrintName(user)
	fmt.Println(name)
}
