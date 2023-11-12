package main

import "fmt"

type data interface{}

type Dog struct {
	Name string
}

func Print(d data) {
	fmt.Println(d)
}

func PrintNum(num interface{}) {
	//if value, ok := num.(int); ok {
	//	fmt.Println("int", value)
	//}
	//if value, ok := num.(string); ok {
	//	fmt.Println("string", value)
	//}

	switch i := num.(type) {
	case int:
		fmt.Println("int", i)
	case string:
		fmt.Println("string", i)
	}

}

func main() {
	d := Dog{"小黑"}

	Print(d)

	//Print(12)
	//Print("123")
	//Print(true)
	//Print([]int{1, 2, 3})
	//Print(make(map[string]string, 2))

	PrintNum(34)
	PrintNum("3455")
	PrintNum(true)
}
