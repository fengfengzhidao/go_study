package main

import (
	"fmt"
	"reflect"
)

func getType(obj any) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("String")
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Struct:
		fmt.Println("Struct")

	}
}

func main() {
	getType("23")
	getType(123)
	getType(struct {
		Name string
	}{})
}
