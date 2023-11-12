package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path, request.UserAgent())
	byteData, err := os.ReadFile("go_study/27.网络编程/http_study/index.html")
	if err != nil {
		writer.Write([]byte("文件不存在"))
		return
	}
	writer.Write(byteData)
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("web server listen addr: http://127.0.0.1:80")
	http.ListenAndServe("127.0.0.1:80", nil)
}
