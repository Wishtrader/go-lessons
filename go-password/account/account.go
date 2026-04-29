package Account 

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890@!~*$#%^&")

type Account struct {
	login string 
	password string 
	url string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *AccountWithTimeStamp) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url, acc.createdAt, acc.updatedAt)
	color.Cyan(acc.login)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN!")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL!")
	}
	newAcc := &AccountWithTimeStamp{
		Account: Account {
			login: login,
			password: password,
			url: urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

