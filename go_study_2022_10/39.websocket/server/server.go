package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connList []*websocket.Conn

// handler
func handler(res http.ResponseWriter, req *http.Request) {
	// 服务升级
	conn, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	connList = append(connList, conn)
	for {
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, c := range connList {
			c.WriteMessage(msgType, []byte(fmt.Sprintf("你说的是：%s吗", msgData)))
		}
		fmt.Println(msgType, string(msgData))

	}
	defer conn.Close()
	fmt.Println("服务关闭")

}

func main() {
	// 回调函数
	http.HandleFunc("/", handler)
	// 绑定服务
	http.ListenAndServe(":8080", nil)
}
