package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var List []string
var mmap map[string]interface{} = map[string]interface{}{}

var wgr = sync.WaitGroup{}

func GetName(funName string) {

	time.Sleep(1 * time.Second)
	mutex.Lock()
	List = append(List, "返回值")
	mmap[funName] = funName + "_返回值"
	mutex.Unlock()
	wgr.Done()

}

func main() {
	fmt.Println(List)
	wgr.Add(2)
	go GetName("nameFunc")
	go GetName("niuName")
	wgr.Wait()

	fmt.Println(List)
	fmt.Println(mmap["nameFunc"])
	fmt.Println(mmap["niuName"])

}
