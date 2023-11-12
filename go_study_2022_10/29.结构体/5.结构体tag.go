package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	Username  string `json:"-"` // - 也不参与序列化
	LookCount int    `json:"look_count"`
	Free      bool   `json:"free"`
	password  string // 小写字母开头的不会参与序列化
}

func main() {
	article := Article{
		Title:     "golang文档",
		Desc:      "golang零基础入门，四小时转岗golang开发",
		Content:   "golang零基础入门，四小时转岗golang开发",
		Username:  "fengfeng",
		LookCount: 1024,
		password:  "ksd8%38",
		Free:      true,
	}

	// 结构体转json
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(jsonData)
	fmt.Println(jsonStr)

}
