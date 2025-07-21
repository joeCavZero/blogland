package api

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/api/database"
	"github.com/joeCavZero/blogland/api/handlers"
)

func SetupAPI(r *mux.Router) {

	database.StartDatabase()

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
