package handlers

import (
	"athelas/usertwist/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func References(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := getToken(request.Header)
	if err != nil {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	if !models.AuthorizedUsers.IsAuthorized(token) {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	var parsedBody ReferencesRequest
	err = json.NewDecoder(request.Body).Decode(&parsedBody)
	if err != nil {
		http.Error(writer, "Bad Request, entity and term keys required", http.StatusBadRequest)
		return
	}

	if parsedBody.Entity != "users" {
		http.Error(writer, "Bad Request", http.StatusTeapot)
		return
	}

	matches, err := models.AuthorizedUsers.Search(parsedBody.Term)
	if err != nil {
		http.Error(writer, "Server Error", http.StatusInternalServerError)
	}

	responseData := make([]ReferenceMatch, 0)

	for _, m := range matches {
		responseData = append(responseData, ReferenceMatch{
			Username: m.Username,
			Data:     m.Password,
		})
	}

	_ = json.NewEncoder(writer).Encode(responseData)
}

func getToken(headers http.Header) (string, error) {
	bearer := headers.Get("Authorization")

	if bearer == "" {
		return "", fmt.Errorf("missing header")
	}

	parts := strings.SplitN(bearer, " ", 2)

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid token")
	}

	return parts[1], nil
}

type ReferencesRequest struct {
	Term   string `json:"term"`
	Entity string `json:"entity"`
}

type ReferenceMatch struct {
	Username string `json:"username"`
	Data     string `json:"data"`
}
