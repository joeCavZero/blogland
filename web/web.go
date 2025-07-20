package web

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/web/routers"
)

func SetupWeb(r *mux.Router) {
	webRouter := r.PathPrefix("/").Subrouter()
	routers.RegisterWebRoutes(webRouter)
}
