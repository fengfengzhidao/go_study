package main

import "fmt"

type MyCode int
type MyAliasCode = int

func (m MyCode) Code() {

}

const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	var age = 1
	if myCode == MyCode(age) {

	}
	if myAliasCode == age {

	}
}
