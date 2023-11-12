package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	IsMan bool
}

func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = tf.Name
		}
		fmt.Println(jsonTag, v.Field(i))
	}
}

func main() {
	s := Student{Name: "枫枫", Age: 24, IsMan: true}
	ParseJson(s)
}
