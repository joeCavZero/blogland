package handlers

import (
	"net/http"

	"github.com/joeCavZero/blogland/web/utils"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetHTMLContentType(w)
	w.WriteHeader(http.StatusOK)
	utils.RenderTemplate(
		w,
		"home_page.html",
		map[string]string{
			"arg": "MACNTOSH Page",
		},
	)
}
