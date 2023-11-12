package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
}

func (User1) Call(name string) {
	fmt.Println("我被调用了", name)
}

func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		if m.Name != "Call" {
			continue
		}
		method := v.Method(i)
		method.Call([]reflect.Value{
			reflect.ValueOf("枫枫"),
		})
	}
}

func main() {
	s := User1{}
	Call(&s)
}
