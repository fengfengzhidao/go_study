package main

import (
	"fmt"
	"go_study/fengfengStudy/init_study"
)

var (
	mysqlPath string
)

func init() {
	fmt.Println("main init1 ...")
}

func init() {
	mysqlPath = "127.0.0.1:3306"
	fmt.Println("main init2 ...")
}

func init() {
	fmt.Println("main init3 ...")
}

func main() {
	fmt.Println(init_study.Name)
	fmt.Println(mysqlPath)
}
