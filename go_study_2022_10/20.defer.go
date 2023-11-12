package main

import "fmt"

func print(n string) {
	fmt.Println(n)
}

func getData(s string) string {
	fmt.Println("getData ", s)
	return s
}

func Defer() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

func main() {

	//fmt.Println("1")
	//
	//defer print("2")
	//defer print("3")
	//
	//fmt.Println("4")

	Defer()

	fmt.Println("o(￣▽￣)ｄ")

}
