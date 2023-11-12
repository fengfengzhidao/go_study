package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
}

func (s Student) Study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}
func (s Student) Info() {
	fmt.Printf("名字：%s 班级：%s\n", s.Name, s.Class.Name)
}

func (s *Student) SetName(name string) {
	s.Name = name
	fmt.Printf("method in : %p\n", s)
}

func main() {
	c1 := Class{
		Name: "三年一班",
	}

	s1 := Student{Class: c1, Name: "枫枫"}
	s1.Study()
	s1.Info()

	s2 := Student{Class: c1, Name: "张三"}
	s2.Study()
	s2.Info()
	fmt.Printf("method out : %p\n", &s2)
	s2.SetName("李四")
	s2.Info()

}
