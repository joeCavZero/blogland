package routers

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/handlers"
	"github.com/joeCavZero/blogland/web/middlewares"
)

func ProfileRouter(r *mux.Router) {
	subRouter := r.PathPrefix("/profile").Subrouter()
	subRouter.HandleFunc("", handlers.ProfileHandler).Methods("GET")
	subRouter.Use(middlewares.AuthMiddleware)
}
