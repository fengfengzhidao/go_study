package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	user := User{Name: "枫枫", Age: 1, Password: "123123"}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}
