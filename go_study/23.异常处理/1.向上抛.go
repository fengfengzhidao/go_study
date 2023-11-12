package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

func server() (res int, err error) {
	res, err = div(2, 1)
	if err != nil {
		// 把错误向上传递
		return
	}

	// 执行其他的逻辑
	res++
	res += 2
	return

}

func main() {
	res, err := server()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
