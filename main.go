package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type account struct {
	login    string
	password string
	URL      string
}

func main() {
	fmt.Println(generatePassword(10))
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

func generatePassword(n int) (str string) {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(chars[rand.IntN(len(chars))])
	}

	str = b.String()
	return str

}
