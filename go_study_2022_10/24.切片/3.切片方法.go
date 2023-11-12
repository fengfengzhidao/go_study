package main

import "fmt"

func Sum(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}

func main() {
	var s []int = []int{1, 2, 3, 4}
	fmt.Println(s, len(s), cap(s))
	s = append(s, 3)
	// 切片是怎么进行扩容的
	fmt.Println(s, len(s), cap(s))
	s = append(s, 4, 5, 6, 7)
	fmt.Println(s, len(s), cap(s))

	// 复制
	var s1 = make([]int, 8, 8)
	copy(s1, s) // 将s复制到s1，复制的个数就是s1的容量大小
	fmt.Println(s1)

	// string与 []byte
	str := "hello 世界"
	fmt.Printf("[]byte(str)=%s,%v\n", []byte(str), []byte(str))
	fmt.Println(len(str))
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}
	fmt.Println(str[6], str[7])

	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 5, 6))
}
