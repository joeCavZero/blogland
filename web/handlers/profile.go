package handlers

import (
	"net/http"

	"github.com/joeCavZero/blogland/web/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetHTMLContentType(w)
	w.WriteHeader(http.StatusOK)
	utils.RenderTemplate(w, "profile.html", nil)
}
