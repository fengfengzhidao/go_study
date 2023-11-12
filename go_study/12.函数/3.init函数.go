package main

import "fmt"

var db int

func init() {
	db = 10
	fmt.Println("连接数据库成功")
	fmt.Println("init-1")
}
func init() {
	fmt.Println("init-2")
}
func init() {
	fmt.Println("init-3")
}

func main() {
	fmt.Println("main", db)
}
