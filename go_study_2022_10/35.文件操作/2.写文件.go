package main

func main() {
	// 如果文件不存在就创建写入
	/*
		file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := file.WriteString("枫枫知道")
		n, err = file.Write([]byte("枫枫知道"))
		fmt.Println(n)
	*/

	// 快速写
	/*
		err := os.WriteFile("w.txt", []byte("你好啊，枫枫！"), 0777)
		fmt.Println(err)
	*/

	// 带缓冲写入
	/*
		file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := bufio.NewWriter(file)
		for i := 0; i < 5; i++ {
			buf.WriteString(fmt.Sprintf("%d %s\n", i+1, "枫枫知道"))
		}
		buf.Flush()
	*/
}
