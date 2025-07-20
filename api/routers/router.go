package routers

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/api/handlers"
)

func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
