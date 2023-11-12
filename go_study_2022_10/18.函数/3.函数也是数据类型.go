package main

import "fmt"

func getRes(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Printf("%v--%T\n", getRes, getRes)

	var setName func(int) string = func(nameCount int) string {
		return "枫枫"
	}
	fmt.Printf("%v--%T\n", setName, setName)

	fmt.Println(setName(2))

	res := func(s string) int {
		return len(s)
	}("hello")
	fmt.Println(res)

}
