package main

import "fmt"

func getAddDiff(a, b int) (int, int) {
	return a + b, a - b
}

func getAddDiffName(a, b int) (sum, diff int) {

	sum = a + b
	diff = a - b

	return a + b, a - b
}

func main() {
	a, b := getAddDiff(2, 1)
	fmt.Println(a, b)

	var y int
	x, y := getAddDiff(2, 1)
	fmt.Println(x, y)

	fmt.Println(getAddDiffName(5, 2))
}
