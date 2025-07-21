package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/api"
	"github.com/joeCavZero/blogland/logger"
	"github.com/joeCavZero/blogland/web"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		logger.ExitErrorf("Error loading .env file: %v", err)
	}

	logger.Infof("Dot Env loaded successfully")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	api.SetupAPI(router)
	web.SetupWeb(router)

	logger.Infof("Starting server on port %s", port)

	err = http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		router,
	)

	if err != nil {
		logger.ExitErrorf("Failed to start server: %v", err)
	}
}
