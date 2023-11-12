package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// SetAge u是拷贝的，修改无效
func (u User) SetAge(age int) {
	u.Age = age
}

// SetName u是执行原结构体User的
func (u *User) SetName(name string) {
	//(*u).Name = name
	u.Name = name
}

func main() {
	user := &User{"张三", 23}
	user.SetName("张四")
	user.SetAge(21) // 这个修改无效
	fmt.Println(user)
}
