package utils

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/joeCavZero/blogland/logger"
)

func RenderTemplate(w http.ResponseWriter, templatePath string, data any) error {
	includesPath := filepath.Join("web", "templates", "includes")
	templatePath = filepath.Join("web", "templates", templatePath)

	includeFiles, err := filepath.Glob(filepath.Join(includesPath, "*.html"))
	if err != nil {
		logger.WebErrorf("Error finding include files: %v", err)
		return err
	}

	allFiles := append([]string{templatePath}, includeFiles...)

	tmpl, err := template.ParseFiles(allFiles...)
	if err != nil {
		logger.WebErrorf("Error executing template: %v", err)
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		logger.WebErrorf("Error executing template: %v", err)
		return err
	}

	return nil
}

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	SetHTMLContentType(w)
	w.WriteHeader(status)
	RenderTemplate(
		w,
		"error.html",
		map[string]any{
			"status_code": status,
		},
	)
}
