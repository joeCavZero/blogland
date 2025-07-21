package routers

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/handlers"
)

func ProtectedRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	return router
}
