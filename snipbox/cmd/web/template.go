package main

import (
	"github/saaicasm/snipbox/internal/models"
	"path/filepath"
	"text/template"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func createTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("internal/ui/html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		files := []string{
			"internal/ui/html/pages/base.tmpl",
			"internal/ui/html/pages/nav.tmpl",
			page,
		}

		ts, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil

}
