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
		http.Error(w, "Erro ao carregar includes", http.StatusInternalServerError)
		logger.WebErrorf("Erro ao carregar includes: %v", err)
		return err
	}

	allFiles := append([]string{templatePath}, includeFiles...)

	tmpl, err := template.ParseFiles(allFiles...)
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		logger.WebErrorf("Erro ao renderizar template: %v", err)
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erro ao executar template", http.StatusInternalServerError)
		logger.WebErrorf("Erro ao executar template: %v", err)
		return err
	}

	return nil
}
