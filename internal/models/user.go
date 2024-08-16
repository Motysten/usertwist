package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Users []User

var AuthorizedUsers = Users{}

func (users *Users) InitUsers(link, username, password, dbName, port string) {
	cfg := mysql.Config{
		Addr:   fmt.Sprintf("%s:%s", link, port),
		User:   username,
		Passwd: password,
		DBName: dbName,
		Net:    "tcp",
	}

	db, connectionErr := sql.Open("mysql", cfg.FormatDSN())
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	result, queryErr := db.Query("SELECT username, password, token FROM users")
	if queryErr != nil {
		log.Fatal(queryErr)
	}

	for result.Next() {
		var user User
		if parseErr := result.Scan(&user.Username, &user.Password, &user.Token); parseErr != nil {
			log.Fatal(parseErr)
		}
		*users = append(*users, user)
		fmt.Println("Added user: ", user.Username)
	}

}

func (users *Users) GetToken(username string, password string) (string, error) {
	for _, u := range *users {
		if u.Username == username && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
			return u.Token, nil
		}
	}

	return "", fmt.Errorf("User not found")
}

func (users *Users) IsAuthorized(token string) bool {
	for _, u := range *users {
		if u.Token == token {
			return true
		}
	}

	return false
}

func (users *Users) Search(term string) (Users, error) {
	matches := make(Users, 0)

	regex, err := regexp.Compile(term)
	if err != nil {
		return Users{}, err
	}

	for _, u := range *users {
		if regex.Match([]byte(u.Username)) {
			matches = append(matches, u)
		}
	}

	return matches, nil
}
