package main

import "fmt"

type NumStr interface {
	Num | Str
}

// ~的意思就是底层数据类型
type Num interface {
	~int | ~int32 | ~int64 | ~uint8
}
type Str interface {
	string
}

type Status uint8

type mySlice1[T NumStr] []T

func main() {
	m1 := mySlice1[int]{1, 2, 3}
	fmt.Println(m1)
	m2 := mySlice1[int64]{1, 2, 3}
	fmt.Println(m2)
	m3 := mySlice1[string]{"hello"}
	fmt.Println(m3)
	m4 := mySlice1[Status]{1, 2, 3}
	fmt.Println(m4)
}
