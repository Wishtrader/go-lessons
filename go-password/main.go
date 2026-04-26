package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login string 
	password string 
	url string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN!")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL!")
	}
	newAcc := &account{
		login: login,
		password: password,
		url: urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@!~*$#%^&")

func main() {
	login := promptData("Enter Your Login")
	password := promptData("Enter Your Password")
	url := promptData("Enter URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid LOGIN or URL! Check Your Input.")
		return
	}
	
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

