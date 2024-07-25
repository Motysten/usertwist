package models

import (
	"athelas/usertwist/internal/cesar"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Users []User

const KEY = 4

var AuthorizedUsers = Users{
	{Username: "user1", Password: cesar.Rotate("passworD", cesar.Right, KEY), Token: "8Mb19OuQAJ"},
	{Username: "user2", Password: cesar.Rotate("passwoRd", cesar.Right, KEY), Token: "1wVtw244Sg"},
	{Username: "admin", Password: cesar.Rotate("admin", cesar.Right, KEY), Token: "a7P47jXfFM"},
	{Username: "login", Password: cesar.Rotate("ADMIN", cesar.Right, KEY), Token: "KlMnOpQrSt"},
}

func (users *Users) GetToken(username string, password string) (string, error) {
	cipheredPassword := cesar.Rotate(password, cesar.Right, KEY)
	for _, u := range *users {
		if u.Username == username && u.Password == cipheredPassword {
			return u.Token, nil

		}
	}

	return "", fmt.Errorf("User not found")
}
