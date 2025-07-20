package api

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/api/database"
	"github.com/joeCavZero/blogland/api/routers"
)

func SetupAPI(r *mux.Router) {

	database.StartDatabase()

	apiRouter := r.PathPrefix("/api").Subrouter()
	routers.RegisterAPIRoutes(apiRouter)

}
