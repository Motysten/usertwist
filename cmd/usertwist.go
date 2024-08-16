package main

import (
	"athelas/usertwist/internal/http/handlers"
	"athelas/usertwist/internal/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port = "8080"
	var dbFQDN string
	var dbPort string
	var dbUser string
	var dbPass string

	initEnvVar("PORT", &port)

	if !initEnvVar("DB_USER", &dbUser) {
		log.Fatal("DB_USER must be set")
	}
	if !initEnvVar("DB_PASS", &dbPass) {
		log.Fatal("DB_PASS must be set")
	}
	if !initEnvVar("DB_LINK", &dbFQDN) {
		log.Fatal("DB_LINK must be set")
	}
	if !initEnvVar("DB_PORT", &dbPort) {
		log.Fatal("DB_PORT must be set")
	}

	fmt.Printf("starting usertwist on port: %s\n", port)

	models.AuthorizedUsers.InitUsers(dbFQDN, dbUser, dbPass, "usertwist", dbPort)

	http.HandleFunc("/", handlers.Banner)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/references", handlers.References)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("usertwist started")
}

func initEnvVar(name string, variable *string) bool {
	if fromEnv := os.Getenv(name); fromEnv != "" {
		*variable = fromEnv
		return true
	}
	return false
}
