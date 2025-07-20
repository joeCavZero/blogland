package routers

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/handlers"
)

func RegisterWebRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(handlers.GetStaticHandler()).Methods("GET")
	router.HandleFunc("/", handlers.HomePageHandler).Methods("GET")
}
