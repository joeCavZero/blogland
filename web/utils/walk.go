package utils

import (
	"github.com/gorilla/mux"
	"github.com/joeCavZero/blogland/logger"
)

func PrintEndPoints(router *mux.Router) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			methods, _ := route.GetMethods()
			logger.Infof("Registered route: %s %s", methods, path)
		} else {
			logger.Infof("Error getting path for route: %v", err)
		}
		return nil
	})
}
