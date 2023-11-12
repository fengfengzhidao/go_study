package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	req := Req{1, 2}
	var res Res
	client.Call("Server.Add", req, &res)
	fmt.Println(res)
}
