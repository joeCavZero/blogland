package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/joeCavZero/blogland/api/database"
	"github.com/joeCavZero/blogland/api/models"
	"github.com/joeCavZero/blogland/api/utils"
	"github.com/joeCavZero/blogland/logger"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var login models.LoginRequest

	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dt := database.New(database.Database)
	ctx := context.Background()

	userFound, err := dt.GetUserByEmailAndPassword(
		ctx,
		database.GetUserByEmailAndPasswordParams{
			Email:    login.Email,
			Password: login.Password,
		},
	)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := models.RandomToken()
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	err = dt.CreateSessionToken(
		ctx,
		database.CreateSessionTokenParams{
			UserID: userFound.ID,
			Token:  token,
		},
	)
	if err != nil {
		http.Error(w, "Failed to create token in database", http.StatusInternalServerError)
		logger.APIErrorf("Failed to create token in database: %v", err)
		return
	}

	response := models.LoginResponse{
		Id:    userFound.ID,
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
