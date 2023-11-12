package main

import (
	"fmt"
	"os"
)

func readDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range dir {
		name := entry.Name()
		nPath := fmt.Sprintf("%s/%s", path, name)
		if entry.IsDir() {
			readDir(nPath)
			continue
		}
		fmt.Println(nPath)

	}
}

func main() {
	readDir("fengfengStudy")
}
