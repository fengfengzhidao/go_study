package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" feng:"name_xxx"`
	Age  int    `json:"age" feng:"age_xxx"`
}

func Print(inter interface{}) {
	switch x := inter.(type) {
	case User:
		x.Name = "张三"
		fmt.Println(x.Name, x.Age)
	}
	//user := inter.(User)
	//user.Name = "xxx"
}

func FPrint(inter interface{}) {
	//t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	//fmt.Println(t.Kind()) // 获取这个接口的底层类型
	//fmt.Println(t.Elem()) // 变量的原始类型
	//fmt.Println(v.Type())
	//fmt.Println(v.Kind())
	//for i := 0; i < t.NumField(); i++ {
	//	//fmt.Println()
	//	// 字段的类型
	//	// 字段名
	//	// 字段的值
	//	// 字段的tag
	//	fmt.Println(
	//		t.Field(i).Type,
	//		t.Field(i).Name,
	//		v.Field(i),
	//		t.Field(i).Tag.Get("feng"),
	//	)
	//}
	e := v.Elem()

	e.FieldByName("Name").SetString("枫枫知道")

	//for i := 0; i < v.NumField(); i++ {
	//
	//}

	//fmt.Println(v.Elem())

}

func PPrint(user *User) {
	user.Name = "王五"
}

func main() {
	user := User{"枫枫", 21}
	//Print(user)
	//fmt.Println(&user)
	FPrint(&user)
	fmt.Println(user)
	//FPrint([]string{"123"})

	//PPrint(&user)
	//fmt.Println(user)
}
