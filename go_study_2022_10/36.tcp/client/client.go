package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, _ := net.DialTCP("tcp", nil,
		&net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 8080})

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Println(string(buf[0:n]))
		}

	}()

	var s string
	for {
		fmt.Scanln(&s)
		if s == "q" {
			break
		}
		conn.Write([]byte(s))
	}

	conn.Close()
}
