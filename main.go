package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	acc.login = "123"
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	var str string
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}

	str = b.String()
	acc.password = str

}

func newAccount(login, password, url string) *account {
	return &account{
		url:      url,
		login:    login,
		password: password,
	}
}

func main() {

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Ввведите URL: ")

	myAccount := new(login, password, url)
	myAccount.generatePassword(12)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
