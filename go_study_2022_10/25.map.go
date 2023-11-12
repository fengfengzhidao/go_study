package main

import "fmt"

func main() {
	// 定义并创建
	var m map[string]string
	// 分配空间
	m = make(map[string]string, 2)
	//fmt.Println(m == nil)
	m["name"] = "zhangsan"

	fmt.Println(m)

	// 第二种创建方式
	m1 := map[string]string{
		//"name": "张三",
	}
	//m1 := map[string]string{}
	m1["age"] = "21"
	fmt.Println(m1)

	// 查找
	// 如果存在这个key，那v就是vlaue，ok就是true，反之v就是对应的空值，ok就是false
	v, ok := m["name1"]
	fmt.Println(v, ok)

	if v, ok = m["name"]; ok {
		fmt.Println("有值：", v)
	} else {
		fmt.Println("没找到！")
	}

	// 删除  没有的也不会报错
	delete(m, "name")

	fmt.Println(m)

	m = map[string]string{
		"name":   "zhansan",
		"addr":   "长沙市",
		"gender": "男",
	}
	fmt.Println(len(m))

	for key, value := range m {
		fmt.Println(key, value)
	}
}
