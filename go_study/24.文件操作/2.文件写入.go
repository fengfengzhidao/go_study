package main

import (
	"fmt"
	"os"
)

func main() {
	//file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_RDWR, 0777)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//byteData, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData))

	//err := os.WriteFile("w1.txt", []byte("你好"), 0666)
	//fmt.Println(err)

	//rFile, err := os.Open("C:\\Users\\26634\\Desktop\\壁纸\\0902\\girl_kimono_umbrella_1027148_1280x720.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer rFile.Close()
	//wFile, err := os.OpenFile("girl.jpg", os.O_CREATE|os.O_WRONLY, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer wFile.Close()
	//io.Copy(wFile, rFile)

	dir, err := os.ReadDir("go_study/24.文件操作")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dir {
		info, _ := entry.Info()

		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}

}
