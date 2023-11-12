package main

import (
	f1 "fmt"
	. "go_study/fengfengStudy/pkg"
	_ "go_study/fengfengStudy/pkg/init_child"
	p "go_study/fengfengStudy/pkg/pkg"
)

func main() {
	f1.Println("main")
	SayHello()
	p.Ya()

}
