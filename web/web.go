package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/handlers"
	"github.com/joeCavZero/blogland/web/routers"
	"github.com/joeCavZero/blogland/web/utils"
)

func SetupWeb(r *mux.Router) {
	subRouter := r.PathPrefix("/").Subrouter()
	subRouter.PathPrefix("/static/").Handler(handlers.GetStaticHandler()).Methods("GET")
	subRouter.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	subRouter.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	subRouter.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	routers.ProfileRouter(subRouter)

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	utils.PrintEndPoints(r)
}
