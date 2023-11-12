package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 创建tcp监听
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 等待连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		addr := conn.RemoteAddr().String()
		fmt.Println(fmt.Sprintf("%s 进来了", addr))
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Println(fmt.Sprintf("%s 出去了", addr))
				break
			}
			fmt.Println(string(buf[0:n]))
			// 原封不动的返回给客户端
			conn.Write(buf[0:n])
		}

	}

	// 获取数据，发送数据
}
