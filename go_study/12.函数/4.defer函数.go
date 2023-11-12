package main

import "fmt"

func main() {
	var age int
	defer fmt.Println("defer2")
	defer func() {
		fmt.Println(age)
	}()

	return
}
