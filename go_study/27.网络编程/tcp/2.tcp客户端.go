package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			// 客户端退出
			if err == io.EOF {
				break
			}
			fmt.Println(string(buf[0:n]))
		}
	}()

	for {
		var text string
		fmt.Printf("请输入你要输入的内容：")
		fmt.Scan(&text)
		if text == "q" {
			conn.Close()
			break
		}
		conn.Write([]byte(text))
	}
}
