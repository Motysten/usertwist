package main

import (
	"athelas/usertwist/internal/http/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port = "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	fmt.Printf("starting usertwist on port: %s\n", port)

	http.HandleFunc("/", handlers.Banner)
	http.HandleFunc("/login", handlers.Login)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("usertwist started")
}
