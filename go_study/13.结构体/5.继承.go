package main

import "fmt"

type Animal struct {
	Name string
}
type Animal1 struct {
	Name1 string
}

type Cat struct {
	Animal
}

func main() {
	var animal = Animal{Name: "猫"}
	cat := Cat{
		Animal: animal,
	}
	fmt.Println(cat.Name)
}
