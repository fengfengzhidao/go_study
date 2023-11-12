package main

import "fmt"

func main() {
	type User struct {
		Name     string
		Age      int
		password string
	}

	type Account struct {
		User
		money float32
		Name  string
	}

	var ac Account = Account{
		money: 21.9,
		User: User{
			Name:     "王五",
			Age:      21,
			password: "978345",
		},
		Name: "账户1",
	}
	fmt.Printf("ac %#v, %T\n", ac, ac)
	var u1 = User{
		Name:     "lisa",
		Age:      22,
		password: "1233fg",
	}
	var ac1 Account = Account{
		money: 21.9,
		User:  u1,
	}
	fmt.Printf("ac1 %#v, %T\n", ac1, ac1)

	fmt.Println(ac.money)
	fmt.Println(ac.Name)
	fmt.Println(ac.User.Name)
	fmt.Println(ac.password)
}
