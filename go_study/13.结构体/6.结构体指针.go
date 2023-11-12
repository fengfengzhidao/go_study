package main

import "fmt"

type UserInfo struct {
	Name string `json:"name"`
}

func (this *UserInfo) SetName(name string) {
	this.Name = name
	fmt.Printf("this:%p\n", this)
}

func main() {
	user := UserInfo{
		Name: "枫枫",
	}
	fmt.Printf("user:%p\n", &user)
	user.SetName("张三")
	fmt.Println(user.Name)
}
