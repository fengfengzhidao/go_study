package main

import "fmt"

type SingInterface interface {
	Sing()
	GetName() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Chicken) GetName() string {
	return c.Name
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Cat) GetName() string {
	return c.Name
}

// sing
func sing(c SingInterface) {
	//ch, ok := c.(Chicken)
	//fmt.Println(ch, ok)

	switch server := c.(type) {
	case Chicken:
		fmt.Println(server)
	case Cat:
		fmt.Println(server)
	default:
		fmt.Println("其他")
	}

	c.Sing()
}

func Print(val any) {
	fmt.Println(val)
}

func main() {
	ch := Chicken{Name: "ik"}
	ca := Cat{Name: "咪咪"}
	sing(ch)
	sing(ca)
	Print(ch)
	Print(ca)
}
