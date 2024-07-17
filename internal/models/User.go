package models

import (
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Users []User

var AuthorizedUsers = Users{
	{Username: "user1", Password: cesar("password1", Right), Token: "8Mb19OuQAJ"},
	{Username: "user2", Password: cesar("password2", Right), Token: "1wVtw244Sg"},
	{Username: "admin", Password: cesar("admin", Right), Token: "a7P47jXfFM"},
	{Username: "login", Password: cesar("password", Right), Token: "KlMnOpQrSt"},
}

func (users *Users) GetToken(username string, password string) (string, error) {
	cipheredPassword := cesar(password, Right)
	for _, u := range *users {
		if u.Username == username && u.Password == cipheredPassword {
			return u.Token, nil

		}
	}

	return "", fmt.Errorf("User not found")
}

const KEY = 4

type Direction int

const (
	Left  Direction = -1
	Right           = 1
)

func cesar(input string, dir Direction) string {

	runes := []rune(input)

	for index, char := range runes {
		switch dir {
		case Left:
			char = char - KEY
		case Right:
			char = char + KEY
		}

		runes[index] = char
	}

	return string(runes)
}
