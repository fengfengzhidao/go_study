package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件复制
	/*
		file1, err := os.ReadFile("xxx.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile("xx1.jpg", file1, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("图片下载成功")
	*/
	// 文件复制

	read, _ := os.Open("xxx.jpg")
	write, _ := os.Create("xx2.jpg")
	n, err := io.Copy(write, read)
	fmt.Println(n, err)

}
