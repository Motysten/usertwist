package handlers

import (
	"athelas/usertwist/internal/models"
	"encoding/json"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var parsedBody LoginRequest
	err := json.NewDecoder(request.Body).Decode(&parsedBody)
	if err != nil {
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return
	}

	token, err := models.AuthorizedUsers.GetToken(parsedBody.Username, parsedBody.Password)
	if err != nil {
		http.Error(writer, "Login or password incorrect", http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(writer).Encode(LoginResponse{
		Username: parsedBody.Username,
		Token:    token,
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
