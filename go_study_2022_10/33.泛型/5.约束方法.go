package main

import (
	"fmt"
	"strconv"
)

type Price int

func (p Price) String() string {
	// int转数字
	return strconv.Itoa(int(p))
}

type Price2 string

func (p Price2) String() string {
	// int转数字
	return string(p)
}

type showPrice interface {
	~int | ~string
	String() string
}

func showPriceFunc[T showPrice](p T) {
	fmt.Println(p.String())

}

func main() {
	var p1 Price = 12
	showPriceFunc(p1)
	var p2 Price2 = "56"
	showPriceFunc(p2)
}
