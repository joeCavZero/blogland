package handlers

import "net/http"

func GetStaticHandler() http.Handler {
	publicHandler := http.FileServer(http.Dir("web/static"))
	return http.StripPrefix("/static/", publicHandler)
}
