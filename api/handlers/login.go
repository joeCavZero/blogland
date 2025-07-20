package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joeCavZero/blogland/api/database"
	"github.com/joeCavZero/blogland/api/models"
	"github.com/joeCavZero/blogland/api/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var login models.LoginRequest

	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// check if the login credentials are in the database
	// For simplicity, we assume the credentials are valid
	var token string
	token, err = database.CreateTokenByCredentials(login.Email, login.Password)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	response := models.AuthResponse{
		Token: token,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

	utils.SetJSONContentType(w)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
