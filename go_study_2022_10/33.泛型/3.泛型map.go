package main

import "fmt"

type myMap[K string | int, V any] map[K]V
type _User struct {
	Name string
}

func main() {
	m1 := myMap[string, string]{
		"key": "fengfeng",
	}
	fmt.Println(m1)
	m2 := myMap[int, _User]{
		0: _User{"枫枫"},
	}
	fmt.Println(m2)
}
