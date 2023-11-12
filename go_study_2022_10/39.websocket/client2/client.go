package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	di := websocket.Dialer{}
	conn, _, err := di.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	go Send(conn)

	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(msgType, string(msgData))
	}
	conn.Close()

}

func Send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}

}
