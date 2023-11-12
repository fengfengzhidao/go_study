package main

import (
	"fmt"
	"reflect"
)

func getValue(obj any) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("String", v.String())
	case reflect.Int:
		fmt.Println("Int", v.Int())
	case reflect.Struct:
		fmt.Println("Struct")

	}
}

func main() {
	getValue("23")
	getValue(123)
	getValue(struct {
		Name string
	}{Name: "枫枫"})
}
