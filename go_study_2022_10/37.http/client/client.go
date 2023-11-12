package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 实例化客户端
	client := new(http.Client)

	// 构造请求对象
	req, _ := http.NewRequest("GET", "http://localhost:8081/index1", nil)
	
	// 发请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ error ] %s\n", err.Error())
		return
	}

	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))

}
