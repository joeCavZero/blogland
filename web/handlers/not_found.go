package handlers

import (
	"net/http"

	"github.com/joeCavZero/blogland/web/utils"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetHTMLContentType(w)
	w.WriteHeader(http.StatusNotFound)
	utils.RenderErrorTemplate(
		w,
		http.StatusNotFound,
	)
}
