package handlers

import "net/http"

func GetStaticHandler() http.Handler {
	publicHandler := http.FileServer(http.Dir("static"))
	return http.StripPrefix("/static/", publicHandler)
}
