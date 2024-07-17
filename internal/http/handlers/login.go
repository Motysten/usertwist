package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Liste des utilisateurs autorisés
var AuthorizedUsers = []User{
	{Login: "user1", Password: "password1", Token: "8Mb19OuQAJ"},
	{Login: "user2", Password: "password2", Token: "1wVtw244Sg"},
	{Login: "toto", Password: "tata", Token: "d7P47HXfFM"},
	{Login: "login", Password: "password", Token: "KlMnOpQrSt"},
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var netUser User
	err := json.NewDecoder(r.Body).Decode(&netUser)
	if err != nil {
		http.Error(w, "Bad raquette", http.StatusBadRequest)
		return
	}

	for _, u := range AuthorizedUsers {
		if u.Login == netUser.Login && u.Password == netUser.Password {
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(w, "Login Accepted here is the token : %v", u.Token)
			return
		}
	}

	http.Error(w, "Login or password incorrect", http.StatusUnauthorized)
}
