package web

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/handlers"
	"github.com/joeCavZero/blogland/web/middlewares"
	"github.com/joeCavZero/blogland/web/routers"
)

func SetupWeb(r *mux.Router) {
	subRouter := r.PathPrefix("/").Subrouter()
	subRouter.PathPrefix("/static/").Handler(handlers.GetStaticHandler()).Methods("GET")
	subRouter.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
	subRouter.HandleFunc("/login", handlers.LoginHandler).Methods("GET")

	protectedRouter := routers.ProtectedRoutes()
	protectedRouter.Use(middlewares.AuthMiddleware)

	subRouter.PathPrefix("/").Handler(protectedRouter)
}
