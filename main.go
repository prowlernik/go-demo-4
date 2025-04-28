package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	URL      string
}

func main() {
	fmt.Println(rand.IntN(10))

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Ввведите URL: ")

	myAccount := account{
		login:    login,
		password: password,
		URL:      url,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	acc.login = "123"
	fmt.Println(acc.login, acc.password, acc.URL)
}

func generatePassword(n int) string {
	rand.
}
