package main

import "fmt"

type Animal interface {
	Sing()
	Dump(time int)
	Rap() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println("Chicken 调用了 Sing方法")
}
func (c Chicken) Dump(time int) {
	fmt.Println("Chicken 调用了 Dump方法")
}
func (c Chicken) Rap() string {
	fmt.Println("Chicken 调用了 Rap方法")
	return "rap"
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println("Cat 调用了 Sing方法")
}
func (c Cat) Dump(time int) {
	fmt.Println("Cat 调用了 Dump方法")
}
func (c Cat) Rap() string {
	fmt.Println("Cat 调用了 Rap方法")
	return "rap"
}

func Sing(animal Animal) {
	animal.Sing()
}

func main() {

	chicken := Chicken{"ik"}
	cat := Cat{"阿狸"}

	Sing(chicken)
	Sing(cat)
}
