package main

import (
	"fmt"
	"soda/password/account"
)

func main() {
	login := promptData("Enter Your Login")
	password := promptData("Enter Your Password")
	url := promptData("Enter URL")

	myAccount, err := Account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid LOGIN or URL! Check Your Input.")
		return
	}
	
	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

