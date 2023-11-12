package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct{}

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func (s Server) Add(req Req, res *Res) error {
	res.Num = req.Num1 + req.Num2
	return nil
}

func main() {
	// 1. 注册rpc服务
	rpc.Register(new(Server))
	// 2. 绑定
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Serve(listen, nil)

}
