package main

import "fmt"

func main() {
	var f float32 = 3.1415926
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)

	num1 := 10
	num2 := 3

	fmt.Printf("%d\n", num1/num2) // 会直接将小数部分抹掉

	num3 := 10.0
	num4 := 3.0
	fmt.Printf("%f\n", num3/num4) // float相除才是float

	//fmt.Printf("%f", num1/num4)  // 类型不同，不能运算
	fmt.Printf("%.2f\n", float64(num1)/num4)
}
