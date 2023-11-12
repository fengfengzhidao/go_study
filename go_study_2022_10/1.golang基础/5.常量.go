package main

import "fmt"

const (
	a int = iota
	b     = 2
	c     = iota
	d     = iota
	e     = 1
	f     = iota
)

const (
	x = iota
	y = iota
)

func main() {

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(x, y)
}
