package main

import (
	"fmt"
	"os"
)

func init() {
	_, err := os.ReadFile("12344")
	if err != nil {
		//log.Fatalln("错误了")
		panic("错误了")
	}

}
func main() {
	fmt.Println("main")
}
