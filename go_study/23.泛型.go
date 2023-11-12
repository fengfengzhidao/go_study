package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func plus[T Number](n1, n2 T) T {
	return n1 + n2
}
func myPrint[T int, K string | int](n1 T, n2 K) {

}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	//plus(1, 2)
	//var u1, u2 = uint(2), uint(3)
	//plus(u1, u2)

	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//user := Response{
	//	Code: 0,
	//	Msg:  "成功",
	//	Data: User{
	//		Name: "枫枫",
	//	},
	//}
	//byteData, _ := json.Marshal(user)
	//fmt.Println(string(byteData))
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "成功",
	//	Data: UserInfo{
	//		Name: "枫枫",
	//		Age:  24,
	//	},
	//}
	//byteData, _ = json.Marshal(userInfo)
	//fmt.Println(string(byteData))

	//var userResponse Response[User]
	//json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫"}}`), &userResponse)
	//fmt.Println(userResponse.Data.Name)
	//var userInfoResponse Response[UserInfo]
	//json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫","age":24}}`), &userInfoResponse)
	//fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)

	//type MySlice[T int|string] []T
	//var mySlice = MySlice[int]{1,2,3}
	//fmt.Println(mySlice[0]+1)

	// map的key只能是基本数据类型
	type MyMap[T string, K any] map[T]K
	var myMap = MyMap[string, string]{
		"name": "12",
	}
	fmt.Println(myMap)

}
